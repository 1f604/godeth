// barebones http server with no external dependencies
// we will use govendor in later tutorials to manage external dependencies as needed
// if you want performance, write your own server or use fasthttp (can reach 280k reqs/s easily)

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

type httpHandler struct {
	cache map[string][]byte
}

func newHTTPHandler() *httpHandler {
	var h httpHandler
	h.cache = make(map[string][]byte)
	return &h
}

// using this instead of reusing a compiled regexp improves reqs/s from 99k to 108k
func isValidFilename(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '.' {
			return false
		}
	}
	return true
}

// serverside caching improves reqs/s from 108k to 167k
func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filename := "index.html"
	var body []byte
	var err error
	requested := path.Base(r.URL.Path)
	if isValidFilename(requested) {
		filename = requested
	}
	if value, exists := h.cache[filename]; exists {
		body = value
	} else {
		body, err = ioutil.ReadFile("dist/" + filename)
		if err == nil {
			h.cache[filename] = body
		} else {
			body = []byte("Error: dist/" + filename + " not found.")
		}
	}
	w.Write(body) //a bit faster than fprintf - brings it from 170k to 180k reqs/sec
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set!!")
	}

	handler := newHTTPHandler()
	//using custom server instead of NewServeMux brings reqs/s from 181k to 183k - no appreciable diff.
	server := http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	log.Fatal(server.ListenAndServe())
}

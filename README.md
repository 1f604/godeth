
# godeth - go development environment test heroku

A simple hello world test program to set up a development environment for go + typescript to work with heroku. Based on [Getting Started with Go on Heroku](https://devcenter.heroku.com/articles/getting-started-with-go).

## Running Locally

Make sure you have node, [Go](http://golang.org/doc/install) and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed. Make sure you have `$GOPATH` set. 

```sh
go get -u github.com/1f604/godeth
cd $GOPATH/src/github.com/1f604/godeth/frontend
npm install tslint typescript parcel-bundler -g
npm install
cd ..
parcel frontend/index.html
go build -o bin/godeth; PORT=8080 ./bin/godeth
```

The app should now be running on [localhost:8080](http://localhost:8080/). Try changing `test.ts` and saving it. You should be able to see the changes in browser immediately. 

You should also install [govendor](https://github.com/kardianos/govendor) since we are going to add dependencies later on. 

## Deploying to Heroku

```sh
heroku create
git wd "initial commit"
git push heroku master
```
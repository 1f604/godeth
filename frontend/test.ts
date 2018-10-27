// tslint makes javascript development less painful
// Build tools like webpack or parcel auto-rebuild your code and refresh your browser on save
// Run this command in terminal to watch for changes: parcel index.html
// For convenience, we commit dist to source control and just serve that from heroku
// Make sure you install tslint before you install the tslint vscode extension
// the alias I added in .git/config lets you simply save, git wd "msg", and git push heroku master
// useful shell aliases: gwm="git wd", gph="git push heroku master"
// workflow: when you're want to see your changes on the server, just do gwd+gph
// you might want to rm -rf dist once in a while. if you do, don't forget to restart parcel
class Startup {
    public static main(): number {
        const x = 2 + "14";
        console.log("Hello World", x);
        return 0;
    }
}

Startup.main();

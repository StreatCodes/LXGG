## LXGG

LXGG is linux containers made easy. LXGG is just a web interface to quickly create, manage and destroy linux containers without having to fiddle with LXC commands. The best part about LXGG is how easy it is to setup, the package `apt install lxgg` (soon..) then goto `localhost:9871` in your web browser to start managing your hosts!

### Building
In lxgg_frontend run `npm install` (requires nodejs)
In lxgg run `go get ./...` requires golang
In root dir `./build.sh`
Then run `./lxgg` in the build dir and lxgg should be running

### Development
Have 2 seperate terminal windows open
In lxgg_frontend run `npm run dev`
In lxgg run `gin` requires `go get github.com/codegangsta/gin`
This setup will auto compile the backend and front end, so you can work on both without having to worry about building / compiling.

### TODO
 - Add remaining fields to create container
 - Add quick action buttons to container list view e.g. start/stop container
 - Improve container info layout
 - Give user ability to customize account settings and logout
 - Store additional info against containers like tags, creator, etc.
 - Create packages for major package managers.
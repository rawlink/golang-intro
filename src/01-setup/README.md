# Setup Your Go Development Environment

In this section, you will setup your Go development environment. Upon completion of this section you should have achieved the following:

* Installed the Go language for your OS.
* Cloned this repository.
* Configured the tutorial as your Go development environment.
* Installed the Glide package manager.
* Installed the IntelliJ Go plugin. If JetBrains products are not your thing, alternatives are provided.
* Opened your first Go project.

## Installing Go

### OSX and Windows

0. Download and install the latest 1.6.x version of Go for your OS from the [Go download page](https://golang.org/dl/).
0. Navigate to the location you downloaded your OS installer to and double-click on the installer.
0. Open a terminal/cmd window and type in `go version` and press enter. You should see output similar to the following:

   `go version go1.6 darwin/amd64`

That’s it! You are done installing Go for your OS.

### Linux

#### Ubuntu

If you are running Ubuntu, the easiest way to install a recent version of Go is through the “Ubuntu containers team” [lxd-stable PPA](https://launchpad.net/~ubuntu-lxc/+archive/ubuntu/lxd-stable).

0. Open a terminal window.
0. Type in `sudo apt-add-repository ppa:ubuntu-lxc/lxd-stable` and press enter.
0. Press enter when prompted to accept adding the lxd-stable repository.
0. Type in `sudo apt-get update` and press enter.
0. Type in `sudo apt-get install golang` and press enter.
0. Open up a terminal/cmd window and type in `go version` and press enter. You should see output similar to the following:

   `go version go1.6 linux/amd64`

That’s it! You are done installing Go for Ubuntu Linux.

#### From the tarball

0. Download the tarball of the 1.6.x version of Go for Linux from the [Go download page](https://golang.org/dl/).
0. Follow the installation instructions for Linux found at the [Go installation documentation page](https://golang.org/doc/install)

## Clone the tutorial repository

Next you will clone this tutorial repository.

0. Open a terminal/cmd window and navigate to the directory from which you would like to clone this repository.
0. Type in `git clone https://github.com/rawlink/golang-intro` and press enter.

## Configure the Go development environment

Now that you have cloned the tutorial repository, you will configure it as your Go development environment. It is assumed that you know how to configure environment variables for your OS. You may need to adjust the examples for your particular OS.

0. Point the GOPATH variable at the cloned tutorial directory. e.g. `GOPATH=/home/rawlink/golang-intro`
0. Set the GO15VENDOREXPERIMENT environment variable to 1. e.g. `GO15VENDOREXPERIMENT=1`
0. Add the $GOPATH/bin directory to your PATH. Don’t worry if this directory doesn’t exist, Go will create it for you. e.g. `PATH=$PATH:$GOPATH/bin`

## Install the Glide package manager

The [Glide](https://github.com/Masterminds/glide) package manager for Go can be found at it’s GitHub repository. Rather than follow the installation instructions contained in the GitHub repository, I encourage you to use the following steps.

0. Open a terminal/cmd window.
0. Type in `go get -u github.com/Masterminds/glide` and press enter. This will retrieve the dependency sources and Glide sources to the source directory, and build them into the pkg and bin directories inside of your $GOPATH.
0. Type in `glide -v` and press enter. You should see output similar to the following:

   `glide version dev`

## Install the IntelliJ IDEA Go plugin

The IDEA Go plugin is compatible with a large set of JetBrains platform IDEs (e.g. IntelliJ IDEA, WebStorm, Android Studio, etc.). It can be found at it's [GitHub page](https://github.com/go-lang-plugin-org/go-lang-idea-plugin).

To install it in IDEA (Other JetBrains platforms may be similar, but no guarantees):

0. Follow the instructions on the [JetBrains site](https://www.jetbrains.com/help/idea/2016.1/managing-enterprise-plugin-repositories.html?origin=old_help) to add the plugin. When you are adding the repository in step 4 of the JetBrains instructions, use the value `https://plugins.jetbrains.com/plugins/alpha/5047` for the weekly build of the plugin.
0. Open the Preferences dialog.
0. Select the Plugins settings.
0. Type `go` in the Plugins search field.
0. Select the Go plugin and install it.

### JetBrains alternatives

For those of you who are not using a JetBrains product, there are several other Go IDEs to choose from.

* Atom with the [go-plus plugin](https://github.com/joefitzgerald/go-plus)
* Sublime with the [GoSublime plugin](https://github.com/DisposaBoy/GoSublime)
* The cross-platform [LiteIDE X](https://github.com/visualfc/liteide)
* [go-vim in docker](https://hub.docker.com/r/mbrt/golang-vim-dev) to ease the pain of setting up and configuring go-vim and friends.
* Others can be found [here](https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins)

## Open your first project

From within IntelliJ:

0. Select File->Open.
0. Navigate to the $GOPATH/src/02-hello directory.
0. Select OK.
0. Open hello.go in the project view.
0. Click on the ‘Setup SDK’ link.
0. Select your Go SDK. If it’s not available select Configure and perform the following steps:
  0. Click on + and select ‘Go SDK’.
  0. Navigate to your installation of the Go SDK. Most of the time, IntelliJ will have automatically selected the correct directory for you.
  0. Select OK in the file selection dialog.
  0. Select OK in the ‘Configure SDK’ dialog.
0. Select OK in the ‘Select Project SDK’ dialog.
0. Click on the ‘Change module type to Go and reload project’ link.
0. Select ‘Reload Project’.

Next: [Hello, Gopher!](../02-hello)

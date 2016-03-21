# Setup

In this section, you will setup your Go development environment. Upon completion of this section you should have achieved the following:

* Installed the Go language for your OS
* Cloned this repository and configured it as your Go development environment
* Installed the Glide package manager
* Installed the IntelliJ Go plugin

## Installing Go

### OSX and Windows

0. Download and install the latest 1.6.x version of Go for your OS from the [Go download page](https://golang.org/dl/).
0. Navigate to the location you downloaded your OS installer to and double-click on the installer.
0. Open up a terminal/cmd window and type in `go version` and press enter. You should see output similar to the following:

   `go version go1.6 darwin/amd64`

That’s it, you are done installing Go for your OS.

### Linux

#### Ubuntu

If you are running Ubuntu, the easiest way to install a recent version of Go is through the “Ubuntu containers team” [lxd-stable PPA](https://launchpad.net/~ubuntu-lxc/+archive/ubuntu/lxd-stable).

0. Open up a terminal window.
0. Type in `sudo apt-add-repository ppa:ubuntu-lxc/lxd-stable` and press enter.
0. Press enter when prompted to accept adding the lxd-stable repository.
0. Type in `sudo apt-get update` and press enter.
0. Type in `sudo apt-get install golang` and press enter.
0. Open up a terminal/cmd window and type in `go version` and press enter. You should see output similar to the following:

   `go version go1.6 linux/amd64`

That’s it, you are done installing Go for Ubuntu Linux.

#### From the tarball

0. Download the tarball of the 1.6.x version of Go for Linux from the [Go download page](https://golang.org/dl/).
0. Follow the installation instructions for Linux found at the [Go installation documentation page](https://golang.org/doc/install)

## Clone this repository

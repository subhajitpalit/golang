# Set-up path 
vi ~/.bash_profile
export JAVA_HOME=$(/System/Library/Frameworks/JavaVM.framework/Versions/CurrentJDK/Home)
alias genc-ms='./genc-ms_darwin_amd64'
export GENC_INSECURE=1
export GENC_TENANT=mstestenv
export GENC_TABLE_STYLE=colored_dark
# below is default cluster
# export GENC_HOST=genesis.paypalcorp.com
# configure below for QA Cluster testing only
#export GENC_HOST=genesis-release.paypalcorp.com
# to support this in your ManagedStage repo, local fork, update below
# grep GENESIS_HOST ple_controller.cfg
# global.GENESIS_HOST=genesis-release.paypalcorp.com
# This is actually your .bashrc file

export GOPATH=$HOME/go-workspace # don't forget to change your path correctly!
export GOROOT=/usr/local/opt/go/libexec
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin



https://sourabhbajaj.com/mac-setup/Go/README.html


Install Golang with Homebrew:
$ brew update
$ brew install golang
When installed, try to run go version to see the installed version of Go.

Setup the workspace:
Add Environment variables:
Go has a different approach of managing code, you'll need to create a single Workspace for all your Go projects. For more information consult : How to write Go Code

First, you'll need to tell Go the location of your workspace.

We'll add some environment variables into shell config. One of does files located at your home directory bash_profile, bashrc or .zshrc (for Oh My Zsh Army)

$ vi .bashrc
Then add those lines to export the required variables

# This is actually your .bashrc file

export GOPATH=$HOME/go-workspace # don't forget to change your path correctly!
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
Create your workspace:
Create the workspace directories tree:

$ mkdir -p $GOPATH $GOPATH/src $GOPATH/pkg $GOPATH/bin
$GOPATH/src : Where your Go projects / programs are located
$GOPATH/pkg : contains every package objects
$GOPATH/bin : The compiled binaries home
Hello world time!
Create a file in your $GOPATH/src, in my case hello.go Hello world program :

package main
import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
Run your first Go program by executing:

$ go run hello.go
You'll see a sweet hello, world stdout

If you wish to compile it and move it to $GOPATH/bin, then run:

 $ go install hello.go
Since we have $GOPATH/bin added to our $PATH, you can run your program from placement :

$ hello
Prints : hello, world

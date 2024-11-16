package main

// package greeting from the github.com/headfirstgo/greeting module
import (
	"fmt"
	"github.com/headfirstgo/greeting"
	"github.com/headfirstgo/greeting/dansk"
	"github.com/headfirstgo/greeting/deutsch"
)

// go tool has another subcommand named 'go get' that can automatically download and install packages.
// to install this repository 'https://github.com/headfirstgo/greeting' we can use the command 'go get github.com/headfirstgo/greeting'
// The go tool will connect to "github.com", download the Git repository at the /headfirstgo/greeting path, and save it
// in the go workspace's 'pkg/mod/github.com/headfirstgo/greeting\@v0.0.0-20190504033635-66e7507184ee/' directory.
// The module name will be added to the go.mod file and the go.mod file will be updated to include the module name.

func main() {
	greeting.Hi()
	greeting.Hello()
	fmt.Println()
	deutsch.Hallo()
	deutsch.GutenTag()
	fmt.Println()
	dansk.Hej()
	dansk.GodMorgen()
}

// To see the package documentation from terminal:
// go doc import-path
// go doc github.com/headfirstgo/greeting
// To see the specific function documentation from terminal:
// go doc github.com/headfirstgo/greeting Hello

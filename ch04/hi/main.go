package main

// We need to import the package before we can use its functions.
// module name (github.com/clebersonp/head-first-go) followed by the path package (/ch04/greeting)
import (
	"fmt"
	"github.com/clebersonp/head-first-go/ch04/greeting"
	"github.com/clebersonp/head-first-go/ch04/greeting/deutsch" // importing path of nested package "greeting/deutsch"
)

func main() {
	// We need the package name and a dot before calls to functions from a different package.
	greeting.Hello()
	greeting.Hi()

	fmt.Println()

	greeting.AllGreetings()

	fmt.Println()

	// calling nested package under the greeting directory
	// we use only the package name that is usually the same name of the directory
	deutsch.Hallo()
	deutsch.GutenTag()
}

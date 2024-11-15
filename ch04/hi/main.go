package main

// We need to import the package before we can use its functions.
// module name (github.com/clebersonp/head-first-go) followed by the path package (/ch04/greeting)
import (
	"fmt"
	"github.com/clebersonp/head-first-go/ch04/greeting"
)

func main() {
	// We need the package name and a dot before calls to functions from a different package.
	greeting.Hello()
	greeting.Hi()

	fmt.Println()

	greeting.AllGreetings()
}

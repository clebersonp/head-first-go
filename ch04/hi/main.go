package main

// We need to import the package before we can use its functions.
import "github.com/clebersonp/head-first-go/ch04/greeting"

func main() {
	// We need the package name and a dot before calls to functions from a different package.
	greeting.Hello()
	greeting.Hi()
}

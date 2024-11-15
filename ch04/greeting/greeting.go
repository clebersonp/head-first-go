// Package greeting isn't "main", it's "greeting"!
package greeting

import "fmt"

// Functions have the first letter capitalized so that functions are exported!

func Hello() {
	fmt.Println("Hello!")
}

func Hi() {
	fmt.Println("Hi!")
}

func AllGreetings() {
	// we would just type the name functions (without the package name qualifier) because we'd be calling the
	// functions from the same package where they're defined.
	Hello()
	Hi()
}

// Package greeting isn't "main", it's "greeting"!
//
// It has some greetings functions on it.
//
// Package comments should begin with "Package" word followed by the package name.
//
// Function comments should begin with the same name of the function they describe.
// We can include code examples in our comments by indenting them.
package greeting

import "fmt"

// Functions have the first letter capitalized so that functions are exported!

// Hello prints the string "Hello!"
func Hello() {
	fmt.Println("Hello!")
}

// Hi prints the string "Hi!"
func Hi() {
	fmt.Println("Hi!")
}

// AllGreetings prints the strings "Hello!" and "Hi!"
func AllGreetings() {
	// we would just type the name functions (without the package name qualifier) because we'd be calling the
	// functions from the same package where they're defined.
	Hello()
	Hi()
}

// To see the package documentation from terminal:
// go doc ./ch04/greeting
// To see the specific function documentation from terminal:
// go doc ./ch04/greeting Hello

// To run on terminal a web server to see the documentation:
// godoc -http=:6060
// Then open a web browser on http://localhost:6060
// To stop the web server press ctrl+c on terminal to stop the godoc server.

// This is the package clause
// This package is a special package
package main

// Any import statements go here
// When we need to import a couple of packages, we can do that within parentheses.
import (
	"fmt"
	"math"
	"strings"
)

// This is the main function. The actual code.
// The "main" function is special. It gets run first when your program runs.
func main() {
	// Calling the fmt package's Println function will be printed to standard output the string "Hello, Go!"
	// The fmt stands for "format". We need to import the package to invoke its functions.
	fmt.Println("Hello, Go!")

	fmt.Println(math.Pi)
	fmt.Println(strings.Title("head first go"))
	fmt.Println("This is scape sequences\n\tnew line and \t\ttab, and \"double quotes\" and backslash \\")
}

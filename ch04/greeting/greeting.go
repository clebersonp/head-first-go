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

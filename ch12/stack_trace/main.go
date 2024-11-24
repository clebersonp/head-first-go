package main

import "fmt"

// When a program panics, all deferred function calls will still be made.
// If there's more than one deferred call, they'll be made in the reverse order.

func main() {
	defer fmt.Println("deferred in main()")
	one()
}
func one() {
	defer fmt.Println("deferred in one()")
	two()
}
func two() {
	defer fmt.Println("deferred in two()")
	three()
}
func three() {
	defer fmt.Println("deferred in three()")
	panic("This call stack's too deep for me!")
}

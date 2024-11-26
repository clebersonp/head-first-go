package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}
func sayBye() {
	fmt.Println("Bye")
}
func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func main() {
	var myFunction func()
	// In a programming language with first-class functions, functions can be assigned to variables, and then called
	// from those variables. To assign a function to a variable, we use the function name without parentheses (call).
	myFunction = sayHi
	// To call a function from a variable, we use the name of the variable followed by parentheses.
	myFunction()

	// First-class functions also allow you to pass functions as arguments to other functions.
	twice(sayHi)
	twice(sayBye)
}

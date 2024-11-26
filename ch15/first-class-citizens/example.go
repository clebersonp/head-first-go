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
func divide(a, b int) float64 {
	return float64(a) / float64(b)
}
func multiply(a, b int) float64 {
	return float64(a * b)
}
func doMath(a, b int, passedFunction func(int, int) float64) float64 {
	return passedFunction(a, b)
}

func main() {
	var myFunction func()
	// In a programming language with first-class functions, functions can be assigned to variables, and then called
	// from those variables. To assign a function to a variable, we use the function name without parentheses (call).
	myFunction = sayHi
	// To call a function from a variable, we use the name of the variable followed by parentheses.
	myFunction()

	// First-class functions also allow you to pass functions as arguments to other functions.
	// Functions type: A function's parameters and return value are part of its type.
	twice(sayHi)
	twice(sayBye)

	// declare a variable of type function that has two parameter of type int and return a float64
	// When we declare a variable of type function, we don't specify the name of parameters, just their types.
	var mathFunction func(int, int) float64
	mathFunction = divide
	fmt.Println(mathFunction(5, 2))

	// Functions that accept a function as a parameter also need to specify the parameters and return types
	// the passed-in function should have.
	// If you try to pass a function with the wrong type, Go will alert you to the problem before your program
	// even compiles.
	fmt.Println(doMath(6, 3, divide))
	fmt.Println(doMath(5, 2, multiply))
}

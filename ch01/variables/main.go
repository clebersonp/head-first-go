package main

import "fmt"

func main() {
	// Declaring variables just use the var keyword followed by the name and the type of the variable
	var quantity int
	var length, width float64
	var customerName string

	// Assigning values to variables
	quantity = 7
	// Assigning multiple variables in a single line
	length, width = 2.5, 3.9
	customerName = "John"

	// Using the variables
	fmt.Println(quantity, length, width, customerName)

	// Declare variable and assign in a single line
	var age int = 30
	fmt.Println(age)

	// We can omit the type if the variable is initialized with a value
	var price = 99.99
	fmt.Println("Price", price)

	// Go zero values
	// Declare a variable without assigning it a value, that variable will contain the 'zero value' for its type.
	// Numeric types: 0
	// Boolean types: false
	// String types: ""
	// Pointer, function, interface, slice, channel, and map types: nil
	var discount float64
	var myBool bool
	var myString string
	fmt.Println("Zero values:", discount, myBool, myString)

	// Short variable declaration
	// Declare a variable and assign a value to it in the same line with ':='
	stock := 100
	isNight := true
	fmt.Println(stock, isNight)

	// Naming rules for variables, functions and types:
	// A name must begin with a letter, and can have any number of additional letters and numbers.
	// If the name begins with a capital letter, it is considered exported and can be accessed from other packages.
	// If it begins with lowercase letter, it is considered unexported and can only be accessed
	// within the current package.
	// For naming conventions we can use camel case.
	// When the meaning of a name is obvious from the context, the Go community's convention is to abbreviate it.
	// to use 'i' instead of 'index' for example.
	// max instead of maximum, and so on.

}

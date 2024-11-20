package main

import "fmt"

// Arrays, slices and maps are no help if you need to mix values of different types.
// The can only be set up to hold values of a single type.
// For that reason, Go provides the 'struct' type to mix different types all together.
// We use the 'struct' key word to define a structure of many types.
// Syntax: struct { field1 string, field2 int, ... }

func main() {
	// declare a struct called myStruct
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	// the struct fields, each set to that type's zero value
	// p%#v will print the value in myStruct as a Go struct literal.
	fmt.Printf("%#v\n", myStruct)

	// Using the 'dot operator' for both assigning values and retrieving them from structs.
	myStruct.number = 3.14
	myStruct.word = "Hello"
	myStruct.toggle = true
	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)
}

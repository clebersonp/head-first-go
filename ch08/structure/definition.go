package main

import "fmt"

// Arrays, slices and maps are no help if you need to mix values of different types.
// The can only be set up to hold values of a single type.
// For that reason, Go provides the 'struct' type to mix different types all together.
// We use the 'struct' key word to define a structure of many types.
// Syntax for declaring struct variable: struct { field1 string, field2 int, ... }

// Syntax for declaring a type definition that's based on an underlying type ( can be any Go underlying type,
// such as array, slice, map, or struct): type myStructure struct { field1 string, field2 bool, ... }
// Just like variables, type definitions can be written within a function (narrow scope) or at the package level (wide scope).

// Don't use an existing type name as a variable name.
// If you've defined a type named car in the current package, and you declare a variable that's also named car,
// the variable name will 'shadow' the type name, making it inaccessible.
// i.e.:
// var car car // shadowing the type car and refers to the type
// var car2 car // refers to the variable, resulting in an error!

// part is the underlying type of struct with these fields
// part is a type definition
type part struct {
	description string
	count       int
}

// car is the underlying type of struct with these fields
// car is a type definition
type car struct {
	name     string
	topSpeed float64
}

// showInfo is a function that takes a part as a parameter and prints its description and count
func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

// minimumOrder is a function that returns a part with a description and count of 100
func minimumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

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

	fmt.Println()

	// declaring and using variables of type 'car' and 'part' both underlying of struct
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Printf("%#v\n", porsche)
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Printf("%#v\n", bolts)
	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)

	fmt.Println()

	// accessing function and passing defined type as a parameter
	showInfo(bolts)

	fmt.Println()
	p := minimumOrder("Another Hex bolts")
	showInfo(p)
}

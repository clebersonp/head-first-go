package main

import "fmt"

// Slices are built on 'top of arrays'.
// Like arrays, slices are made up of multiple elements, all the same type.
// Unlike arrays, functions are available that allow us to add extra elements onto the end of a slice.

func main() {
	// Unlike arrays, slices are dynamically-sized and can grow or shrink.
	// We don't specify the length of a slice when we create it.
	// Unlike with array variable, declaring a slice variable doesn't automatically create a slice.
	var mySlice []string
	fmt.Printf("Length: %d, Data: %#v, Nil: %t\n", len(mySlice), mySlice, mySlice == nil)

	// To create a slice with built-in make function.
	mySecondSlice := make([]string, 7)
	fmt.Printf("Length: %d, Data: %#v, Nil: %t\n", len(mySecondSlice), mySecondSlice, mySecondSlice == nil)
	mySecondSlice[0] = "do"
	mySecondSlice[1] = "re"
	mySecondSlice[2] = "mi"
	fmt.Println(mySecondSlice)
	for i, note := range mySecondSlice {
		fmt.Println(i, note)
	}

	// Initialize the slice using a 'slice literal'
	// There's no need to use the make function to initialize a slice literal.
	// Using a slice literal in the code will create the slice and prepopulate it.
	numbers := []int{1, 9, 86}
	fmt.Println(numbers)

}

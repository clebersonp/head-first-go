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

	// Slice is only a way of viewing an underlying array.
	// When you take a slice of an underlying array, you can only "see" the portion of the array's elements that
	// are visible through the slice.

	// Create a slice by using the 'slice operator [n:n]' in either on array or another slice.
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(underlyingArray)
	// start at index 0 and end at index 3 but not including index 3 (2)
	slice1 := underlyingArray[0:3]
	fmt.Println(slice1)

	// Slice operator has defaults for both the start and stop indexes.
	// If you don't specify the start index, it will default to 0.
	fmt.Println(underlyingArray[:2])

	// If you omit the stop index, it will default to the length of the underlying array.
	fmt.Println(underlyingArray[2:])

	// If you omit both the start and stop indexes, it will default to the entire array.
	fmt.Println(underlyingArray[:])

	// Many slices can share the same underlying array and can even overlap.
	slice2 := underlyingArray[0:3] // Output: [a b c]
	slice3 := underlyingArray[2:5] // Output: [c d e]
	// The letter c or at index 2 appears in both slices
	fmt.Println(slice2, slice3)

	// Change the underlying array will change all the slices.
	underlyingArray[2] = "X"
	fmt.Println(underlyingArray)
	fmt.Println(slice2, slice3)

	// It's recommended to avoid creating slices with slice operator from arrays.
	// It's recommended to create slices using the make function or slice literal instead.
	// Go will create an underlying array to hold the slice elements automatically.
	// This way, you never have to work with the underlying array.

	// Using the built-in 'append' function to add elements to the end of that slice.
	// It returns a new, larger slice with all the same elements as the original slice, plus the new elements
	// added onto the end.
	slice := []string{"a", "b", "c"}
	// We need to return the new slice back to the same slice variable we passed to append to avoid some inconsistencies
	slice = append(slice, "d")
	slice = append(slice, "e", "f")
	slice4 := []string{"g", "h", "i"}
	// special operator '...(spread operator)' adds all the elements of the slice4 to the end of the slice
	slice = append(slice, slice4...)
	fmt.Println(slice)

	// As with arrays, slice that no value has been assigned to, will get the 'zero value' for that type back:
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice, boolSlice)

	// Unlike arrays, the slice variable itself also has a zero value: it's 'nil'.
	var intSlice []int
	var stringSlices []string
	fmt.Printf("intSlice: %#v, stringSlices: %#v\n", intSlice, stringSlices)

	// Functions in Go are intentionally flexible and written to treat a 'nil' slice as an empty slice.
	// In other languages we need to check if a 'slice' is 'null' before we can use it. Not in Go.
	fmt.Printf("Capacity: %d, Length: %d, Data: %#v, Nil: %t\n", cap(intSlice), len(intSlice), intSlice, intSlice == nil)
	intSlice = append(intSlice, 1, 2, 3)
	fmt.Printf("Capacity: %d, Length: %d, Data: %#v, Nil: %t\n", cap(intSlice), len(intSlice), intSlice, intSlice == nil)

}

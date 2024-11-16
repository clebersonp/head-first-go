package main

import "fmt"

func main() {

	// Lopping over arrays safely with 'for...range loop'
	// In the range form, you provide a variable that will hold the integer index of each element,
	// another variable that will hold the value of the element itself, and the array you want to loop over.
	// for index, value := range myArray { loop block here. }
	// Because it's safer and easier to read, you'll see the for loop's range form used most often when working
	// with arrays and other collections.

	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	for index, note := range notes {
		fmt.Println(index, note)
	}

	fmt.Println()

	// Using the blank identifier (_) to ignore the index with 'for...range loop'
	// If we stop using the index variable from our for...range loop, we'll get a compile error.
	// And the same would be tru if we didn't use the variable that holds the element value.
	for _, note := range notes {
		fmt.Println(note)
	}

	fmt.Println()

	// black identifier for array element value
	for index, _ := range notes {
		fmt.Println(index)
	}

}

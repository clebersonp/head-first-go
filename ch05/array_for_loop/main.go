package main

import "fmt"

func main() {
	// Trying to access an index that is outside the array will cause a 'panic', an error that occurs while your program
	// is running. Run time error: index out of range
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	for i := 0; i <= 2; i++ {
		fmt.Println(i, notes[i])
	}

	fmt.Println()

	// We can use the built-in 'len' function to get the length of an array.
	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}
}

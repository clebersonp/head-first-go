package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Variadic functions are functions that can take any number of arguments
// Variadic must be the last argument of a function
// We can pass any number of arguments to a variadic function
// We can pass a slice to a variadic function using the 'spread operator ...'
// Variadic arguments are passed to the function as a slice
func average(numbers ...float64) float64 {
	// create a variable to hold the sum of the numbers
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	// Read command-line arguments
	arguments := os.Args[1:]
	var numbers []float64
	// iterate over the command-line arguments
	for _, argument := range arguments {
		// try to convert each command-line argument to a float
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	// print the sum and average of the numbers from command-line interface
	// passing a slice to a variadic function using the 'spread operator ...'
	fmt.Printf("Numbers: %#v. Length: %d. Average: %0.2f\n", arguments, len(arguments), average(numbers...))
}

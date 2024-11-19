package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Read command-line arguments
	arguments := os.Args[1:]
	// create a variable to hold the sum of the numbers
	var sum float64
	// iterate over the command-line arguments
	for _, argument := range arguments {
		// try to convert each command-line argument to a float
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		// add the converted number to the sum
		sum += number
	}
	// count the number of arguments
	sampleCount := float64(len(arguments))
	// print the sum and average of the numbers from command-line interface
	fmt.Printf("Numbers: %#v. Length: %d. Total: %0.2f. Average: %0.2f\n", arguments, len(arguments), sum, sum/sampleCount)
}

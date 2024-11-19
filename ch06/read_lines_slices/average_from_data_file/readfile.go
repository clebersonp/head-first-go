package main

import (
	"fmt"
	"github.com/clebersonp/head-first-go/ch06/read_lines_slices/datafile"
	"log"
)

// Run the program at terminal inside the ch05/average_from_data_file directory:
// go run readfile.go
// If we 'go install' the program, the binary file will be moved to the go workspace's /bin directory in the
// same directory as the binary file.

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Length: %d. Numbers: %v. Total: %0.2f. Average: %0.2f\n", len(numbers), numbers, sum, sum/sampleCount)
}

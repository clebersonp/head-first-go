package main

import "fmt"

func main() {

	numbers := [3]float64{71.8, 56.2, 89.5}
	calculateAverage(numbers)
}

func calculateAverage(numbers [3]float64) {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Total: %0.2f. Average: %0.2f", sum, sum/sampleCount)
}

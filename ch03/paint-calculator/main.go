package main

import (
	"fmt"
	"log"
)

// variable declaration package level:
var metersPerLiter float64

// function declaration:
// func keyword followed by function name, parameters and return type (optional)
// functions can return multiple values or none
// parameters is a list of parameters separated by comma and each parameter has a name and type
// function is a block of code that performs a specific task that have one or more statements
// function parameters and variables declared within a function block are only in scope within that function block.
// Variables declared at the package level (outside any function) are visible in any function in the package.
func paintNeeded(width, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	if metersPerLiter <= 0 {
		return 0, fmt.Errorf("meters per liter of %0.2f is invalid", metersPerLiter)
	}
	area := width * height
	return area / metersPerLiter, nil
}

func main() {
	// accessing variable at package level
	metersPerLiter = 10.0
	var total float64
	// call function of same package
	// the variable value will be copied to the function parameter. This is called 'pass-by-value'.
	// function parameters receive a copy of the arguments from the function call.
	amount, err := paintNeeded(4.2, 3.0)
	// always handle errors!
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	amount, err = paintNeeded(5.2, 3.5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f liters\n", total)
}

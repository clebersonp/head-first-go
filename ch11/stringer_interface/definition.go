package main

import "fmt"

// The Stringer interface:
// It has a method called String() that returns a string to allow any type to decide how it will be displayed when printed.
// Any type is a fmt.Stringer if it has a String() method that returns a string.

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

type Milliliters float64

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}

func main() {
	gallons := Gallons(4.0)
	milliliters := Milliliters(2000.0)
	liters := Liters(1.5)
	// Many functions in the fmt package check whether the values passed to them satisfy the Stringer interface,
	// and call their String() methods if so.
	// fmt.Println will call the String() method for any type that implements the Stringer interface
	fmt.Println(gallons)
	fmt.Println(gallons.String()) // It has the same effect as fmt.Println(gallons) because it calls the String() method
	fmt.Printf("%s\n", liters)
	fmt.Println(liters.String())
	fmt.Print(milliliters, "\n")
	fmt.Println(milliliters.String())
}

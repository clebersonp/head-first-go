package main

import (
	"fmt"
	"reflect"
)

// Go defined types most often use structs as their underlying types, but they can also be based on ints, strings,
// booleans, floats, or any other type.

// Liters is a defined type that's based on a float64 (underlying float type)
type Liters float64

// Gallons is a defined type that's based on a float64 (underlying float type)
type Gallons float64

func main() {
	var carFuel Gallons
	var busFuel Liters
	fmt.Println("carFuel:", carFuel, ", busFuel:", busFuel)
	// We can do a conversion to a defined type from any value of the underlying type.
	// Write the type you want to convert to, followed by the value you want to convert in parentheses, using type conversion.
	carFuel = Gallons(10.0) // convert a float64 to Gallons
	busFuel = Liters(240.0) // convert a float64 to Liters
	fmt.Println("carFuel:", carFuel, ", busFuel:", busFuel)
	fmt.Printf("carFuel type: %T, busFuel type: %T\n", carFuel, busFuel)
	fmt.Println(
		"carFuel underlying type:",
		reflect.TypeOf(carFuel).Kind(),
		", busFuel underlying type:",
		reflect.TypeOf(busFuel).Kind(),
	)

	// I can't assign a value of a defined type to another defined type, even if the other type has the same underlying type.
	// This helps protect against errors.
	// carFuel = Liters(10.0) // compile error
	// But I can convert between types that have the same underlying type.
	// Liters can be converted to Gallons and vice versa, because both have an underlying type float64.
	// These are legal, but incorrect, because 40.0 Liters is not 40.0 Gallons. It's only for study purposes.
	carFuel = Gallons(Liters(40.0))
	busFuel = Liters(Gallons(63.0))
	fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel)
	// A way to convert with the correct value is to multiply the value by factor of conversion.
	carFuel = Gallons(Liters(40.0) * 0.264) // Liters to Gallons
	busFuel = Liters(Gallons(63.0) * 3.785) // Gallons to Liters
	fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel)

}

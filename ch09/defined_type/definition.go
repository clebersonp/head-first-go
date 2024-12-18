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

// Milliliters is a defined type that's based on a float64 (underlying float type)
type Milliliters float64

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

	// Defined type supports all the same operations as the underlying type.
	fmt.Println(Liters(1.2) + Liters(3.4))
	fmt.Println(Gallons(4.5) == 4.5)
	type Title string
	fmt.Println(Title("Hello") + " world!")

	// Defined types cannot be used in operations together with values of a different type, even if the other type
	// has the same underlying type.
	//fmt.Println(Liters(1.2) + Gallons(4.5)) // compile error
	//fmt.Println(Title("Hi") == 2.0) // compile error
	// If I want to make operations with a defined type and a value of a different type, I need to convert the value.
	fmt.Println(Liters(3.4) + Liters(Gallons(4.5)))

	// Using functions to convert between defined types, Gallons to Liters and vice versa.
	fmt.Printf("Car fuel: %0.1f gallons\n", LitersToGallons(Liters(40.0)))
	fmt.Printf("Bus fuel: %0.1f liters\n", GallonsToLiters(Gallons(30.0)))

}

// The Go language is simplified by 'not' supporting overloading of functions.
// So, in Go we cannot have functions with the same name, even the number and types of arguments are different.
// To avoid creating many different functions we can use methods instead.

func LitersToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func MillilitersToGallons(m Milliliters) Gallons {
	return Gallons(m * 0.000264)
}

func GallonsToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func GallonsToMilliliters(g Gallons) Milliliters {
	return Milliliters(g * 3785.41)
}

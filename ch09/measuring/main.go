package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

// This way we can use the same method name for different types
// For these methods, we don't need a receiver type because we're not modifying the receivers, and the values
// don't consume much memory, so it's fine for the parameter to receive a copy of the value instead of a pointer.

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l * 1000.0)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func (m Milliliters) ToLiters() Liters {
	return Liters(m * 0.001)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func main() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())
	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", milk, milk.ToMilliliters())
}

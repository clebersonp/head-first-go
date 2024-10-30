package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Convert a value from one type to another type.
	// Provide the type you want to convert a value to, immediately followed by the value
	// you want to convert in parentheses.
	var myInt int = 42
	fmt.Println(myInt, reflect.TypeOf(myInt))
	myFloat64 := float64(myInt) + 5.83
	fmt.Println(myFloat64, reflect.TypeOf(myFloat64))
	fmt.Println(myFloat64, reflect.TypeOf(int(myFloat64)), int(myFloat64))
}

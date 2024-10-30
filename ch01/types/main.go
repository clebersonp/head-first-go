package main

import (
	"fmt"
	"reflect"
)

// Go is statically typed, which means  that it knows the types of your values are even before your program runs.
// We can use reflect package to get the type of variable at runtime.

func main() {
	fmt.Println(reflect.TypeOf(42))           // int
	fmt.Println(reflect.TypeOf(3.1415))       // float64
	fmt.Println(reflect.TypeOf(true))         // bool
	fmt.Println(reflect.TypeOf("Hello, Go!")) // string
}

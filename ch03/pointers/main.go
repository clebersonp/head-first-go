package main

import (
	"fmt"
	"reflect"
)

// Go is a pass-by-value language, meaning that function parameters receive a copy of any arguments from the caller.
// Passing pointer to a function can save memory and improve performance depends on the size of struct that will be passed.
// Values that represent the address of a variable are known as pointers, because they point to the location
// where the variable can be found. e.g.: 0x1040a128 -> 3.1415 (float64)
// We use & to get the address of a variable and * to the type of pointer.
// We can use the 'reflect.TypeOf' function to show us the types of pointers.
// When we can change a value through a pointer we call it 'pass-by-reference'.
// We can pass a pointer to a function parameter to change the original value.
func double(number *int) {
	// showing the type of pointer number with %T (formatting verbs) and package 'reflect'
	fmt.Printf("Type of number: %T, Type of number: %v, Pointer address: %v\n", number, reflect.TypeOf(number), number)
	// accessing variable through pointer *number (value at number address | dereference)
	*number *= 2
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func main() {
	amount := 6
	// passing a pointer to a function (& 'address of operator')
	double(&amount)
	fmt.Println(amount)
	truth := true
	negate(&truth)
	fmt.Printf("Truth: %t\n", truth)
	lies := false
	negate(&lies)
	fmt.Printf("Lies: %t\n", lies)
}

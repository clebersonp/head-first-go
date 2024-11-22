package main

import "fmt"

// When we need to change some value of a type, we need to use a pointer receiver.
// That is necessary because Go is a pass-by-value language (copy value).
// So, we need pass-by-reference (copy value of reference) instead.
// Syntax: func (r *receiverType) methodName(anyParameter type) returnType {}

type Number int

func (n *Number) Double() {
	// n is a pointer and to change the value of the variable, we need to use '* operator' to access the value at n address.
	// We call the '* operator' of deference to access the value at n address.
	*n *= 2
}

type MyType string

// A method with value receiver
func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

// A method with pointer receiver
func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func main() {
	number := Number(4)
	fmt.Println("Original value of number:", number)
	// The variable number is not a pointer of Number type,
	// Go will automatically pass the address of number to the Double method because it's a pointer receiver.
	number.Double()
	fmt.Println("number after calling Double:", number)

	// We can call the Double method on a pointer of Number type as well.
	numberPointer := &number
	numberPointer.Double()
	fmt.Println("numberPointer after calling Double:", *numberPointer) // '* operator' deference to access the value

	// Go automatically converts the value to pointer and vice versa.
	// We can call methods on both value and pointer types.
	// The code works fine. But it's not recommended.
	// All methods of a type should have either value or pointer receiver, not both.
	// These examples are just for demonstration purposes.
	value := MyType("a value")
	pointer := &value
	value.method()
	value.pointerMethod()
	pointer.method()
	pointer.pointerMethod()

	// There are two limitations of pointer receiver:

	// We can only get pointers to values that are stored in variables.
	// This not compile:
	// &value.pointerMethod()

	// Go can automatically convert values to pointers, but only if the receiver value is stored in a variable.
	// This not compile:
	// MyType("a value").pointerMethod()

}

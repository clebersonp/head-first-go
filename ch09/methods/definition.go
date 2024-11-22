package main

import "fmt"

// A method is very similar to a function.
// The only difference is that the first parameter of a method is a receiver parameter, in parentheses before
// the function name.
// Syntax: func (r receiverType) methodName(anyParameter type) returnType {}
// Receiver type can be any underlying type, such as a struct or basic types.
// We use the receiver type to refer to the object that the method is called on. Very similar to 'this' in Java.
// If we need to change the value of some field of the receiver, we need to use a pointer receiver.
// Syntax: func (r *receiverType) methodName(anyParameter type) returnType {}
// To call a method on a value, we use dot notation.
// Syntax: valueOfType.methodName(anyParameter)
// The method that are defined in the same package as the type becomes associated with all values of that type.
// Aside from the fact that they're called on a receiver, methods are otherwise pretty similar to any other function.

// MyType is a type with an underlying type of string
type MyType string

// sayHi is a method of MyType type.
// To be considered a MyType method, it must be defined with the same receiver type as the type definition.
// The receiver name is not important, but the receiver type is.
// The receiver type determine which type will be associated with the method.
// By convention, the first letter of the receiver's type name in lowercase.
func (m MyType) sayHi() {
	// as m is the underlying type og string, we can use it as a string. So, we use it to print
	// the value of receiver parameter.
	// so, receiver parameter is more than another parameter, we can use it to access any other
	// fields, methods, and so on.
	// It's like a 'this' in Java
	fmt.Println("Hi,", m)
}

func main() {
	// create a MyType value
	value := MyType("a MyType value")
	// call sayHi on that value with dot notation.
	value.sayHi()
	anotherValue := MyType("another MyType value")
	anotherValue.sayHi()
}

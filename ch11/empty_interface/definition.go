package main

import (
	"fmt"
	"strings"
)

// Empty interface is an interface that has no methods.
// Go has two ways to declare an empty interface with built-in types:
// interface{} or any.
// any is an alias for interface{}: type any = interface{}
// We can create a custom empty interface by declaring a type with no methods.
// i.e.: type MyEmptyInterface interface{}
// We can create our own alias for empty interface.
// i.e.: type MyEmptyInterface = interface{} or type MyEmptyInterface = any
// If we declared an interface type that didn't require any methods at all, It would be satisfied by any type. It would
// be satisfied by 'all' types.
// Empty interface is used to accept values of 'any' type.
// The empty interface doesn't require any methods to satisfy it, and so it's satisfied by 'all' type.
// If you have a value with the empty interface as its type, there's not much you can do with it.
// There are no methods you can call on a value with the empty interface type!
// To call methods on a value with the empty interface type, you need to convert it to a concrete type with 'type assertion'.

type myAnyAlias = interface{}

type Month string

func (m Month) Upper() string {
	return strings.ToUpper(string(m))
}

func AcceptAnything(thing myAnyAlias) {
	fmt.Printf("Type: %[1]T, value: %#[1]v\n", thing)
	// safe type assertion conversion to Month type
	month, ok := thing.(Month)
	if ok {
		fmt.Printf("Month: %s\n", month.Upper())
	}
}

func testConcreteType(m Month) {
	fmt.Println(m)
}

func main() {
	AcceptAnything(3.1415)
	AcceptAnything(14)
	AcceptAnything("A string")
	AcceptAnything(true)
	AcceptAnything(Month("November"))

	// When you have a variable of an interface type,
	// the only methods you can call on it are those defined in the interface.
	var i any
	i = Month("April")
	// testConcreteType(i) // It's not compiled
	testConcreteType(i.(Month))
}

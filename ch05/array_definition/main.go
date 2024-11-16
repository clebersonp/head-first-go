package main

import "fmt"

func main() {
	// An array can hold a collection of elements of the same type.
	// An array holds a specific number of elements, and it cannot grow or shrink.
	// To declare a variable that holds an array, you need to specify the number of elements it holds in square brackets.
	// ([]), followed by the type of elements the array holds.

	// To declare an array, we need to specify the number of elements it will hold and its type.
	// [n] number of elements array will hold
	// string is the type of elements the array will hold
	var myArray [4]string
	fmt.Println(myArray)

	// To set the array elements' values or retrieve values later, we need to use an index.
	// The index of an array starts with 0 until the number of elements minus 1.
	myArray[0] = "Hello"
	myArray[1] = "World"
	myArray[2] = "!"
	myArray[3] = "!"
	// The length of an array
	fmt.Println(len(myArray))
	fmt.Println(myArray[0])
	fmt.Println(myArray[1])
	fmt.Println(myArray[2])
	fmt.Println(myArray[3])
	fmt.Println(myArray)

	// Declare an array of 5 integers
	var primes [5]int
	// Assign values to the array by the index.
	primes[0] = 2
	primes[1] = 3
	// Print the first element of the primes array
	fmt.Println(primes[0])

	// Zero values in arrays
	// As with variables, when an array is created, all the values it contains are initialized to the 'zero value' for
	// the type that array holds. So an array of int values is filled with zeros by default:
	// 0 (Zero) for numbers, false for booleans, "" for strings
	// nil for pointers, functions, interfaces, slices, channels, and maps.
	fmt.Println(primes[2])
	fmt.Println(primes[3])
	fmt.Println(primes[4])

	// Zero values can make it safe to manipulate an array even if you haven't explicitly assigned a value to it.
	// When a array is created, all the values it contains are initialized to the zero value for the type the array holds.
	var active [3]bool
	fmt.Println(active[0])
	fmt.Println(active[1])
	fmt.Println(active[2])

	// Initialized an array using 'array literal'.
	var notes = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes[3], notes[6], notes[0])
	// we don't need to specify all the elements at 'array literal'
	// we can declare an array using 'short variable declaration'.
	odd := [5]int{1, 3, 5, 7}
	fmt.Println(odd)

	// We can spread 'array literals' over multiple lines.
	text := [3]string{
		"This is a series of long string",
		"which would be awkward to place",
		"together on a single line.",
	}
	// Println function in the fmt package know how to handle arrays
	fmt.Println(text)

	// Using verb '%#v' to print the array
	// When formatted by '%#v', arrays appear in the result as Go array literals.
	fmt.Printf("%#v\n", text)

}

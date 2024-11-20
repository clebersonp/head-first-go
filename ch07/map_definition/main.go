package main

import "fmt"

// A map is a collection of key-value pairs
// Keys are an easy and fast way to get data back out of a map
// The map can use any type as a key (as long as values of that type can be compared (Comparable) using ==)
// All the keys must be the same type.
// All values must be the same type.
// But a key don't have to be the same type as the value
// map literals: myMap := map[string]float64{"a": 1.2, "b": 3.4}
// myMap := map[string]float64{} creates an empty map

func main() {
	var isPrime map[int]bool
	isPrime = make(map[int]bool)
	isPrime[4] = false
	isPrime[5] = true
	fmt.Println("Primes:", isPrime)
	fmt.Println("4 is prime?", isPrime[4])
	fmt.Println("5 is prime?", isPrime[5])

	ages := map[string]int{"Alice": 31, "Bob": 29}
	fmt.Println(ages)

	// Empty map with make function
	workers := make(map[string]int)
	fmt.Println("Workers:", workers)

	// Empty map with map literal
	employees := map[string]int{}
	fmt.Println("Employees:", employees)

	// map will be 'nil' if it is not initialized
	// If we try to put a new key-value pair into a nil map, it will panic
	var myMap map[string]int
	//myMap["key"] = 1 // panic: assignment to entry in nil map
	fmt.Printf("myMap: %#v, it is nil: %t\n", myMap, myMap == nil)

	// Zero values within maps
	// As with arrays and slices, if you access a map key that hasn't been assigned to, you'll get the zero value back.
	numbers := make(map[string]int)
	numbers["one"] = 1
	fmt.Printf("%#v\n", numbers["one"])
	fmt.Printf("%#v\n", numbers["two"])

}

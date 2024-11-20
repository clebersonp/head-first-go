package main

import (
	"fmt"
	// import path to import other package, magazine in this case
	// All exported names from the magazine package are available. It'll can be struct, function, constants, methods and so on.
	"github.com/clebersonp/head-first-go/ch08/magazine"
)

func main() {
	// We can access exported names from the magazine package because its name is capitalized (exported)
	var s magazine.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)

	// We can create a variable using 'struct literal'
	// This is a literal for a Subscriber struct.
	subscriber := magazine.Subscriber{
		Name:   "Aman Singh",
		Rate:   4.99,
		Active: true,
	}
	fmt.Printf("%#v\n", subscriber)

	// We can omit one or more fields in a struct literal.
	// In that case, all omitted fields will have their zero value
	newSubscriber := magazine.Subscriber{
		Rate: 7.90,
	}
	fmt.Printf("%#v\n", newSubscriber)

	// Creating an employee
	joy := magazine.Employee{
		Name:   "Joy Carr",
		Salary: 60000,
	}
	fmt.Printf("Employee: %#v\n", joy)

	var otherSubscriber magazine.Subscriber
	// all field will be set to their zero value, even the fields of type struct
	fmt.Printf("Subscriber: %#v\n", otherSubscriber)
	// As 'Go' initialize all fields with zero value, we can access to the inner structs' fields without problem.
	fmt.Printf("otherSubscriber.HomeAddress.City: %#v\n", otherSubscriber.HomeAddress.City)
	// we call this 'otherSubscriber.HomeAddress.PostalCode' of 'dot operator chaining'
	otherSubscriber.HomeAddress.PostalCode = "67200"
	fmt.Printf("otherSubscriber.HomeAddress.PostalCode: %#v\n", otherSubscriber.HomeAddress.PostalCode)
}

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

	fmt.Println()

	// create and initialize a struct with 'struct literal'
	anna := magazine.Employee{
		Name:   "Anna Smith",
		Salary: 80_000,
		HomeAddress: magazine.Address{
			Street:     "123 Main St",
			City:       "New York",
			State:      "NY",
			PostalCode: "12345",
		},
		// Permission is an anonymous struct or embedded struct, we can access its fields directly or via Permission
		Permission: magazine.Permission{
			CanRead:  true,
			CanWrite: true,
		},
	}
	fmt.Printf("Employee: %#v\n", anna)
	// accessing the inner struct Permission fields
	fmt.Printf("anna.Permission.CanRead: %#v\n", anna.Permission.CanRead)
	fmt.Printf("anna.Permission.CanWrite: %#v\n", anna.Permission.CanWrite)
	// accessing the inner struct Address fields directly because it's an embedded struct
	// This is similar to 'java inheritance', but Go has no inheritance
	fmt.Printf("anna.CanRead: %#v\n", anna.CanRead)
	fmt.Printf("anna.CanWrite: %#v\n", anna.CanWrite)
}

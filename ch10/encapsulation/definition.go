package main

import (
	"fmt"
	"github.com/clebersonp/head-first-go/ch10/calendar"
	"log"
)

// Encapsulation: a way to protect your struct type's fields from that invalid data.

// We need encapsulate validation data to setter methods that will be used to validate data.
// Setter methods are methods used to set fields or other values within a defined type's underlying value.
// By convention, Go setter methods are usually named in the form SetX, where X is the name of the field.

func main() {
	// don't compile because the fields are unexported and the struct is another package
	//date := calendar.Date{year: 2019, month: 5, day: 27}
	//fmt.Printf("%#v\n", date)

	// Using setter methods
	date := calendar.Date{}
	// date is a type value, Go automatically converts to a pointer to call the method that has a pointer receiver
	if err := date.SetYear(2020); err != nil {
		log.Fatal(err)
	}
	if err := date.SetMonth(6); err != nil {
		log.Fatal(err)
	}
	if err := date.SetDay(15); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", date)
}

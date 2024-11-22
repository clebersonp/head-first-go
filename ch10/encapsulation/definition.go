package main

import (
	"fmt"
	"github.com/clebersonp/head-first-go/ch10/calendar"
	"log"
)

// Encapsulation: a way to protect your struct type's fields from that invalid data.
// Go developers generally only rely on encapsulation when it's necessary, such as when field data needs to be
// validated by setter methods.

// We need encapsulate validation data to setter methods that will be used to validate data.
// Setter methods are methods used to set fields or other values within a defined type's underlying value.
// By convention, Go setter methods are usually named in the form SetX, where X is the name of the field.
// Unexported variables, struct fields, functions, and methods can still be accessed by exported functions and methods
// in the same package.

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
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())

	// Create a calendar event with embedded Date type
	event := calendar.Event{}
	// We can call setter methods on the embedded Date type with 'dot notation' in Date type value
	if err := event.Date.SetYear(2024); err != nil {
		log.Fatal(err)
	}
	// We can call setter methods on the embedded Date type with 'dot notation' directly on the event type value
	if err := event.SetMonth(8); err != nil {
		log.Fatal(err)
	}
	if err := event.SetDay(10); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d-%.2d-%d\n", event.Date.Year(), event.Month(), event.Day())
}

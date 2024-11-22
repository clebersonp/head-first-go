package main

import "fmt"

// Encapsulation: a way to protect your struct type's fields from that invalid data.

// We need encapsulate validation data to setter methods that will be used to validate data.
// Setter methods are methods used to set fields or other values within a defined type's underlying value.
// By convention, Go setter methods are usually named in the form SetX, where X is the name of the field.

type Date struct {
	// reduce the visibility of fields to protect them from invalid data
	// This will enforce that the field must be set using a setter method
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) {
	// Go automatically converts to Date value, so we don't need to use '* operator''
	d.year = year
}

func (d *Date) SetMonth(month int) {
	d.month = month
}

func (d *Date) SetDay(day int) {
	d.day = day
}

func main() {
	date := Date{year: 2019, month: 5, day: 27}
	fmt.Printf("%#v\n", date)

	// Using setter methods
	date = Date{}
	// date is a type value, Go automatically converts to a pointer to call the method that has a pointer receiver
	date.SetYear(2020)
	date.SetMonth(6)
	date.SetDay(19)
	fmt.Printf("%#v\n", date)
}

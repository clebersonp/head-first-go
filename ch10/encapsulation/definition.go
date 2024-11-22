package main

import "fmt"

// Encapsulation: a way to protect your struct type's fields from that invalid data.

type Date struct {
	Year  int
	Month int
	Day   int
}

func main() {
	date := Date{Year: 2019, Month: 5, Day: 27}
	fmt.Printf("%#v\n", date)
}

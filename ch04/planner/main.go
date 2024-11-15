package main

import (
	"fmt"
	"github.com/clebersonp/head-first-go/ch04/dates"
)

func main() {
	days := 3
	fmt.Println("Your appointment is in", days, "days")
	// call and use constants from another package
	fmt.Println("With a follow-up in", days+dates.DaysInWeek, "days")

	// calling a function from another package
	fmt.Println("With a follow-up in", dates.WeeksToDays(days), "days")

}

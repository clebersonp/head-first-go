package main

import "fmt"

func main() {

	// Any percent signs (%) will be treated as the start of a 'formatting verb', a section of the string that
	// will be substituted with a value in a particular format. The remaining arguments are used as values
	// with those verbs. The letter following the percent sign indicates which verb to use.
	// %s for string values, %d for integer values, %f for float values.
	// All those are 'formatting verbs'.
	fmt.Printf("The %s cost %d cents each.\n", "apple", 23)
	fmt.Printf("That will be $%0.2f please.\n", 0.23*5)
	phraseWithDayOfWeek := fmt.Sprintf("Today is %s", "Sunday")
	fmt.Println(phraseWithDayOfWeek)
	fmt.Println()

	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 15)
	fmt.Printf("A string: %s\n", "hello")
	fmt.Printf("A boolean: %t\n", false)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n")
	fmt.Printf("%v %v %v\n", "", "\t", "\n")
	fmt.Printf("%#v %#v %#v\n", "", "\t", "\n") // with %#v, we can actually see them

	fmt.Println()

	// Formatting with padding
	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("----------------------------------------")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)

	fmt.Println()

	// Formatting fractional numbers widths
	// %5.3f -> 5 is the minimum width for entire number, 3 is the minimum width for the fractional part,
	// f means formatting verb type
	fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
	fmt.Printf("%%.1f: %.1f\n", 12.3456)
	fmt.Printf("%%.2f: %.2f\n", 12.3456)

}

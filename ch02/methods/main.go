package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# R#cks!"
	replacer := strings.NewReplacer("#", "o")
	// Methods are special functions that are associated with values of a particular type.
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}

package main

import "fmt"

func main() {

	// String literals are written surrounded by double quotation marks ("),
	// but runes literals are written with single quotation marks (').
	// Go uses the Unicode standard for storing runes. Runes are kept as numeric codes,
	// not the characters themselves.
	fmt.Println('A', 'B', '\t', '\n', '\\', ' ')

}

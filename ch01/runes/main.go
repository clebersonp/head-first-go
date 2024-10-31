package main

import (
	"fmt"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"log"
	"unicode"
)

func main() {

	// String literals are written surrounded by double quotation marks ("),
	// but runes literals are written with single quotation marks (').
	// Go uses the Unicode standard for storing runes. Runes are kept as numeric codes,
	// not the characters themselves.
	fmt.Println('A', 'B', '\t', '\n', '\\', ' ')

	str := "São Paulo, Conceição, Água, D'água, Amanhã ŷ ý û, acentuación De dónde eres, excepción"
	result, err := removeAccents(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func removeAccents(str string) (string, error) {
	// Transforms string with runes package from golang.org/x/text dependency
	tChain := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, err := transform.String(tChain, str)
	return result, err
}

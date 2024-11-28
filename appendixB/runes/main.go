package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	asciiString := "ABCDE"
	utf8String := "АБВГД"

	// len function returns the number of bytes in the string
	// asciiString has 5 bytes, 1 byte for each character. Each rune requires 1 byte
	fmt.Println(len(asciiString))
	// utf8String has 10 bytes, 2 bytes for each character. Each rune requires 2 bytes to store.
	fmt.Println(len(utf8String))

	// The length of a string in 'characters' is not the same as the length of a string in 'bytes'.
	// If you want the length of a string in characters, use the unicode/utf8 package's RuneCountInString function.
	// This function will return the correct number of characters in the string.
	fmt.Println(utf8.RuneCountInString(asciiString)) // 5 runes
	fmt.Println(utf8.RuneCountInString(utf8String))  // 5 runes

	// Don't work with partial strings with bytes.
	asciiBytes := []byte(asciiString)
	utf8Bytes := []byte(utf8String)
	asciiBytesPartial := asciiBytes[3:]
	utf8BytesPartial := utf8Bytes[3:]
	fmt.Println(string(asciiBytesPartial))
	fmt.Println("wrong:", string(utf8BytesPartial)) // wrong

	// Working with partial strings safely means converting the string to runes, not bytes.
	// To work with partial string, you should convert them to a slice of 'rune' values rather than a slice of byte values.
	asciiRunes := []rune(asciiString)
	utf8Runes := []rune(utf8String)
	asciiRunesPartial := asciiRunes[3:]
	utf8RunesPartial := utf8Runes[3:]
	fmt.Println(string(asciiRunesPartial))
	fmt.Println("correct:", string(utf8RunesPartial))

	// that's ok because ascii characters only need 1 byte
	for index, currentByte := range asciiBytes {
		fmt.Printf("%d: %s\n", index, string(currentByte))
	}

	// Fails because utf8 russian characters need 2 bytes
	for index, currentByte := range utf8Bytes {
		fmt.Printf("%d: %s\n", index, string(currentByte))
	}

	// safe approach using rune on 'range loop'
	for position, currentRune := range asciiString {
		fmt.Printf("%d: %s\n", position, string(currentRune))
	}
	for position, currentRune := range utf8String {
		fmt.Printf("%d: %s\n", position, string(currentRune))
	}

	// Remember: Anytime you want to work with just part of a string , convert it do runes, not bytes!

}

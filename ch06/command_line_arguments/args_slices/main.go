package main

import (
	"fmt"
	"os"
)

// Run the program at terminal:
// go run args_slices/main.go 78.2 65.39 12.54

// Getting command-line arguments from the os.Args slice

func main() {
	// os.Args is a slice of strings, the command-line arguments
	// The first argument is the name of the program
	// From second argument on, there are the real command-line arguments
	fmt.Println("All arguments:", os.Args)

	// To print only the command-line arguments without the program name using the 'slice operator [:]'
	fmt.Println("Only program arguments:", os.Args[1:])
}

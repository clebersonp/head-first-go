package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Run the program at terminal:
// go run readfile.go

func main() {
	// open the data file for reading
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	// close the file to free resources as soon as possible when the main functions return.
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	// create a new scanner for the file
	scanner := bufio.NewScanner(file)
	// loops until the end of the file
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// It's possible that the bufio.Scanner encountered an error while scanning through the file.
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

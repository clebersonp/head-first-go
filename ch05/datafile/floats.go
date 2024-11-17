package datafile

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([3]float64, error) {
	// Go arrays are fixed in size. They can't grow or shrink.
	// But the data.txt file can have as many lines as the user wants to add.
	// we are using hardcode size of the array.
	// If we put more elements to the file, the program will panic: index out of range
	// This is the array limitation
	var numbers [3]float64
	// open the data file for reading
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
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
	// This variable will track which array index we should assign to.
	i := 0
	// loops until the end of the file
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	// It's possible that the bufio.Scanner encountered an error while scanning through the file.
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}

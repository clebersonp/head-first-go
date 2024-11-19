package datafile

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([]float64, error) {
	// Using slices to manage the data reading from the file
	var numbers []float64
	// open the data file for reading
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
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
		var number float64
		number, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	// It's possible that the bufio.Scanner encountered an error while scanning through the file.
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

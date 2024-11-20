package datafile

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// GetStrings returns a slice of strings from a file
func GetStrings(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not open the file %s: %w", fileName, err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

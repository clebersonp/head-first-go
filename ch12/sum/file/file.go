package file

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

func GetFloats(fileName string) ([]float64, error) {
	fl, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	var numbers []float64
	scanner := bufio.NewScanner(fl)
	for scanner.Scan() {
		var number float64
		number, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	CloseFile(fl)
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

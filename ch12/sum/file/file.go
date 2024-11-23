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
	// Recovering from errors using deferred functions calls
	// defer (delay) will execute the CloseFile function before the GetFloats function returns.
	// It'll be the last thing to execute.
	// defer execute even any error occurs.
	// The 'defer' keyword ensures a function call takes place, even if the calling function exits early.
	// Important: We need to put the call function with defer immediately after the call to 'OpenFile'.
	defer CloseFile(fl)
	var numbers []float64
	scanner := bufio.NewScanner(fl)
	for scanner.Scan() {
		var number float64
		number, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err // Now, even if an error is returned here, CloseFile will still be called, because 'defer' keyword.
		}
		numbers = append(numbers, number)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err() // Now, even if an error is returned here, CloseFile will still be called, because 'defer' keyword.
	}
	return numbers, nil // The 'CloseFile' will still be called, because 'defer' keyword and it'll be the last thing to execute.
}

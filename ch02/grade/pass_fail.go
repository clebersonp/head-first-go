// pass_fail reports whether a grade is passing or failing. These comments will be ignored by the compiler.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	// multiple return values from method
	// we can use blank identifier as a placeholder for the error value to ignore the error return value
	//input, _ := reader.ReadString('\n')
	// we can handle the error
	input, err := reader.ReadString('\n')
	// conditionally we check if there's an error before logging it and kill the program.
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)

	fmt.Println()
	size, err := fileSize("ch02/grade/pass_fail.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pass_fail.go size:", size)
}

func fileSize(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

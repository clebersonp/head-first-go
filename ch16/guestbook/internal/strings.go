package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetSignatures(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

func AddSignature(fileName, signature string) error {
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	// 0644 is in octal notation. Any series of digits preceded by a 0 Zero will be treated as an octal number.
	// Go lets you write number using octal notation.
	file, err := os.OpenFile(fileName, options, os.FileMode(0644))
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(file, signature)
	if err != nil {
		return err
	}
	err = file.Close()
	return err
}

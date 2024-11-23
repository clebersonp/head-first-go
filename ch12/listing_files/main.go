package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// list files in a directory
func listFiles(path string) error {
	fmt.Println(path)
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, f := range files {
		fullPath := filepath.Join(path, f.Name())
		if f.IsDir() {
			// recursively call the function while the current file is a directory
			// passing the full directory's path to the function
			// 'recursion' allows a function to call itself.
			if err = listFiles(fullPath); err != nil {
				return err
			}
		} else {
			fmt.Println(fullPath)
		}
	}
	return nil
}

func main() {
	if err := listFiles("ch12/listing_files/my_directory"); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// list files in a directory
func listFiles(path string) {
	fmt.Println(path)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Returning error from listFiles(\"%s\") call\n", path)
		// changing the error handling to 'panic' because panic crashes the program and stops execution.
		// error handling with recursion will return all the stack execution before return the listFiles function.
		// All the recursive calls to listFiles exit.
		// Calling 'panic' may simplify the code, but it also crashes the program.
		// That doesn't seem like much of an improvement.
		panic(err)
	}
	for _, f := range files {
		fullPath := filepath.Join(path, f.Name())
		if f.IsDir() {
			// recursively call the function while the current file is a directory
			// passing the full directory's path to the function
			// 'recursion' allows a function to call itself.
			listFiles(fullPath)
		} else {
			fmt.Println(fullPath)
		}
	}
}

func main() {
	listFiles("ch12/listing_files/my_directory")
}

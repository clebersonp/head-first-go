package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 3 keywords to use before a program crashes:
// defer
// panic
// recover

// We only can recover from a panic, and we use a deferred function call to recover that panic.

func reportPanic() {
	// The 'recover' function returns whatever value was originally passed to the 'panic' function.
	p := recover()
	// nothing happens if p is nil
	if p == nil {
		return
	}
	if err, ok := p.(error); ok {
		fmt.Println("Recovered from panic:", err)
	} else {
		// That's ok, we're waiting for an error type.
		// Everything else will just panic again.
		// Most programs should panic only in the event of an unanticipated error. This is a good situation for panic.
		panic(p)
	}
}

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
		// Generally, calling 'panic' should be reserved for 'impossible' situations: errors that indicate a bug in the
		// program, not a mistake on the user's part.
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
	// calling a method with defer before the potential panic
	// If a deferred function calls the built-in 'recover' function, the program will recover from a panic state (if any).
	// When a program panics, any deferred function call will still be made, allowing cleanup code to be executed
	// before a crash.
	defer reportPanic()
	// panic("some other issue") // to test the reinstating a panic with a string message instead of an error type
	listFiles("ch12/listing_files/my_directory")
}

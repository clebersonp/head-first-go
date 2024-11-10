package main

import (
	"fmt"
	"math/rand" // it's an import path
	"strconv"
	"time"
)

// rand is a package name and rand.New is a function from rand package
// By convention, the last (or only) segment of the import path is also used as the package name.
// So if the import path is "archive", the package name will be "archive", and if the import path
// is "archive/zip", the package name will be "zip".
// To get a random number we need to seed the random number generator with a different value each time.
// Otherwise, the random will always generate the same number
var random = rand.New(rand.NewSource(time.Now().Unix()))

func getUserInput() (int, error) {
	var inputUser string
	_, err := fmt.Scanln(&inputUser)
	if err != nil {
		return 0, err
	}
	// Convert the input string to an integer
	return strconv.Atoi(inputUser)
}

func randomNumber(end int) int {
	// Generate a random number between 1 and 100
	return random.Intn(end) + 1
}

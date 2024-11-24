package main

import (
	"fmt"
	"time"
)

// Goroutines let your program work on several different tasks at once. Your goroutines can coordinate their work
// using channels, which let them send data to each other and synchronize so that one goroutine doesn't get ahead
// of another.
// Goroutines let you take full advantage of computers with multiple processors,
// so that your programs run as fast as possible!

func main() {
	// we initialize a new goroutine with the 'go' keyword
	// the goroutine will run concurrently with the main goroutine
	// The order is not guaranteed
	go func() {
		fmt.Print("Hello ")
	}()
	go func() {
		fmt.Print("World! ")
	}()
	// This is not the best approach.
	// Channels or using package sync is the best one.
	// Just for demonstration purposes.
	time.Sleep(200 * time.Millisecond)
	fmt.Println("\ndone!")
}

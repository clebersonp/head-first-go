package main

import (
	"fmt"
	"time"
)

func sendLetters(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "a"
	time.Sleep(1 * time.Second)
	channel <- "b"
	time.Sleep(1 * time.Second)
	channel <- "c"
	time.Sleep(1 * time.Second)
	channel <- "d"
}

func main() {
	start := time.Now()
	fmt.Println(start)
	channel := make(chan string)
	go sendLetters(channel)
	// When the main goroutine wakes up, it receives four values from the channel.
	// The sendLetters goroutine was blocked until the first value was received,
	// so now we have to wait while later values are sent.
	time.Sleep(5 * time.Second)
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	end := time.Now()
	fmt.Println(end)
	fmt.Println("Duration:", end.Sub(start))
}

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
	channel := make(chan string, 4)
	go sendLetters(channel)
	// When the main goroutine wakes up, it receives four values from the channel.
	// The sendLetters goroutine only will block after filling the buffer, that is, after sending four values
	// so now we have to wait 5 seconds to receive all four values from the channel buffered.
	// So when the main goroutines wakes up after 5 seconds, it immediately receives the four values.
	// This allows the program to complete in only 5 seconds!
	time.Sleep(5 * time.Second)
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	// In a buffered channel, mustn't call a received operation on an empty channel, it will cause a panic 'deadlock'
	// because it will wait forever for a value to be sent to the channel but no one will send a value
	//fmt.Println(<-channel, time.Now())
	end := time.Now()
	fmt.Println(end)
	fmt.Println("Duration:", end.Sub(start))
}

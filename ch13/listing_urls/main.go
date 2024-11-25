package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// In Go, concurrent tasks are called goroutines. Other programming languages hava similar concept called threads,
// but goroutines require less computer memory than threads, and less time to start up and stop, meaning you can
// run more goroutines at once.
// Goroutines allow for concurrency: pausing one task to work on others.
// An in some situations they allow parallelism: working on multiple tasks simultaneously!
// The 'main' function of every Go program is started using a goroutine, so every Go program runs at least one goroutine.

// We can't use function return values in a go statement.
// Go won't let your use the return value from a function called with a go statement, because there's no guarantee
// the return value will be ready before we attempt to use it.
// The way to communicate between goroutines is using channels.
// Channels block operation in the current goroutines. A send operation block the sending goroutine until another
// goroutine executes a receiver operation on the same channel. And vice versa: a receiver operation blocks the
// receiving goroutine until another goroutine executes a send operation on the same channel. This behavior allows
// goroutines to 'synchronize' their actions.
// A receiving goroutine waits until another goroutine sends a value.
// To declare a variable that holds a channel, you use the 'chan' keyword, followed by the type of values that
// channel will cary.
// Syntax: var myChannel chan type
// To create a channel we need to call the built-in 'make' function.
// Syntax: var myChannel chan float64 -> Declare a variable to hold a channel.
//         myChannel = make(chan float64) -> Actually create the channel.
// Or: myChannel := make(chan float64) -> Create a channel and declare a variable at once.
// To send a value to channel:
// Syntax: myChannel <- value
// To receive a value from channel:
// Syntax: value := <- myChannel

// We don't directly control when goroutines run. A goroutine run in a different order each time the program is run.
// But if the order of goroutines run in is important, we need to synchronize them using channels.

// Syntax: go myFunction()
// 		'go' keyword | Function call

type sizeResult struct {
	url      string
	size     int
	duration time.Duration
}

func main() {
	start := time.Now()
	results := make(chan sizeResult)
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://golang.org/doc",
		"https://google.com",
		"https://github.com",
		"https://stackoverflow.com",
		"https://x.com",
	}
	for _, url := range urls {
		go responseSize(url, results)
	}
	// wait for all goroutines to finish
	// The number of goroutines is equal to the number of urls
	// I can't range over the channel directly because nobody will close it. So, It'll block forever for waiting
	// anyone to close the channel
	for range urls {
		result := <-results
		fmt.Printf("Size %7d in %13s for %s\n", result.size, result.duration, result.url)
	}
	end := time.Now().Sub(start)
	fmt.Println("Duration:", end)
}

// channel is a way to communicate between goroutines and this way of communication is only to send a value to the channel
func responseSize(url string, channel chan<- sizeResult) {
	start := time.Now()
	fmt.Printf("Getting %s\n", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	end := time.Now().Sub(start)
	channel <- sizeResult{url: url, size: len(body), duration: end}
}

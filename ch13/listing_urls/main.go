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

func main() {
	start := time.Now()
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
		go responseSize(url)
	}
	time.Sleep(time.Second)
	end := time.Now().Sub(start)
	fmt.Println("Duration:", end)
}

func responseSize(url string) {
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
	fmt.Printf("Size %d for %s\n", len(body), url)
}

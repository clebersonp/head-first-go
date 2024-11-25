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

// We don't directly control when goroutines run. A goroutine run in a different order each time the program is run.
// But if the order of goroutines run in is important, we need to synchronize them using channels.

// Syntax: go myFunction()
// 		'go' keyword | Function call

func main() {
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

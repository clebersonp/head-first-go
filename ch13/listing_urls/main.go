package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://golang.org/doc",
	}
	for _, url := range urls {
		responseSize(url)
	}
}

func responseSize(url string) {
	fmt.Printf("Getting %s ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Size %d\n", len(body))
}

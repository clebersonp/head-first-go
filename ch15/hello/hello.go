package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, request *http.Request, message string) {
	log.Printf("Received a request. User-Agent: %#v, Host: %s\n", request.Header.Get("User-Agent"), request.Host)
	_, err := io.WriteString(writer, message)
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, request, "Hello, web!")
}
func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, request, "Salut, web!")
}
func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, request, "Namaste, web!")
}

func main() {
	// Go supports 'firs-class functions', which allow you to pass functions to other functions.
	// The first-class function will be store to be called later when a matching request path is received.
	// In a programming language with first-class functions, functions can be assigned to variables, and then called
	// from those variables. To assign a function to a variable, we use the function name without parentheses (call).
	// Its handlers request to the path "/hello" (resource path)
	// Any time a request with a path of "/hello" is received, the function englishHandler will be called
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	fmt.Println("Web server is running at http://localhost:8080")
	// http.ListenAndServe will starts up the web server, and it'll run forever, unless it encounters an error.
	// ":8080" is short for "localhost:8080"
	log.Fatal(http.ListenAndServe(":8080", nil))
}

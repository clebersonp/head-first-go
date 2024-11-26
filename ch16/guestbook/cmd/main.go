package main

import (
	"io"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(w, "Hello World!")
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

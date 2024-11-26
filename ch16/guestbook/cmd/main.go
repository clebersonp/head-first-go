package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	root      = "/"
	guestbook = "/guestbook"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func redirectViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Redirecting to", guestbook)
	http.Redirect(w, r, guestbook, http.StatusSeeOther)
}

func viewHandler(w http.ResponseWriter, _ *http.Request) {
	// The html/template package is based on the text/template package.
	// The html/template has some extra security features needed for working with HTML.
	html, err := template.ParseFiles("ch16/guestbook/resources/view.html")
	check(err)
	err = html.Execute(w, nil)
	check(err)
}

func main() {
	http.HandleFunc(root, redirectViewHandler)
	http.HandleFunc(guestbook, viewHandler)
	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

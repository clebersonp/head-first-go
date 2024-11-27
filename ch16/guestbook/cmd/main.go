package main

import (
	"github.com/clebersonp/head-first-go/ch16/guestbook/internal"
	"html/template"
	"log"
	"net/http"
)

const (
	root            = "/"
	guestbook       = "/guestbook"
	newGuestbook    = "/guestbook/new"
	createGuestbook = "/guestbook/create"
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
	html, err := template.ParseFiles("ch16/guestbook/resources/pages/view.html")
	check(err)
	signatures, err := internal.GetStrings("ch16/guestbook/resources/data/signatures.txt")
	log.Printf("Signatures: %#v\n", signatures)
	check(err)
	err = html.Execute(w, internal.Guestbook{SignatureCount: len(signatures), Signatures: signatures})
	check(err)
}

func newHandler(w http.ResponseWriter, _ *http.Request) {
	html, err := template.ParseFiles("ch16/guestbook/resources/pages/new.html")
	check(err)
	err = html.Execute(w, nil)
	check(err)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Adding signature")
	log.Println(r.FormValue("signature"))
	http.Redirect(w, r, guestbook, http.StatusSeeOther)
}

func main() {
	http.HandleFunc(root, redirectViewHandler)
	http.HandleFunc(guestbook, viewHandler)
	http.HandleFunc(newGuestbook, newHandler)
	http.HandleFunc(createGuestbook, createHandler)
	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

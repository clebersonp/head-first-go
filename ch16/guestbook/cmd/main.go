package main

import (
	"github.com/clebersonp/head-first-go/ch16/guestbook/internal"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// Variables
const (
	addr              = ":8080"
	signatureDataForm = "signature"
)

// Files
const (
	viewHtmlFilePath  = "ch16/guestbook/resources/pages/view.html"
	newHtmlFilePath   = "ch16/guestbook/resources/pages/new.html"
	signatureFilePath = "ch16/guestbook/resources/data/signatures.txt"
)

// Routes
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
	http.Redirect(w, r, guestbook, http.StatusFound)
}

func viewHandler(w http.ResponseWriter, _ *http.Request) {
	// The html/template package is based on the text/template package.
	// The html/template has some extra security features needed for working with HTML.
	html, err := template.ParseFiles(viewHtmlFilePath)
	check(err)
	signatures, err := internal.GetSignatures(signatureFilePath)
	log.Printf("Signatures: %#v\n", signatures)
	check(err)
	err = html.Execute(w, internal.Guestbook{SignatureCount: len(signatures), Signatures: signatures})
	check(err)
}

func newHandler(w http.ResponseWriter, _ *http.Request) {
	html, err := template.ParseFiles(newHtmlFilePath)
	check(err)
	err = html.Execute(w, nil)
	check(err)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Adding signature")
	signature := strings.TrimSpace(r.FormValue(signatureDataForm))
	if signature == "" {
		http.Redirect(w, r, newGuestbook, http.StatusFound)
		return
	}
	err := internal.AddSignature(signatureFilePath, signature)
	check(err)
	http.Redirect(w, r, guestbook, http.StatusFound)
}

func main() {
	http.HandleFunc(root, redirectViewHandler)
	http.HandleFunc(guestbook, viewHandler)
	http.HandleFunc(newGuestbook, newHandler)
	http.HandleFunc(createGuestbook, createHandler)
	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(addr, nil))
}

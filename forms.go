// forms.go
package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type contactDetails struct {
	Email   string
	Subject string
	Message string
}

var details contactDetails

func info(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "_info_\n")
	fmt.Fprintf(w, "Email: %s\n", details.Email)
	fmt.Fprintf(w, "Subject: %s\n", details.Subject)
	fmt.Fprintf(w, "Message: %s\n", details.Message)
}

func reader(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details.Email = r.FormValue("email")
	details.Subject = r.FormValue("subject")
	details.Message = r.FormValue("message")

	// do something with details
	// _ = details

	// tmpl.Execute(w, struct{ Success bool }{true})
	fmt.Fprintf(w, "_thanks_\n")
}

func gorilla(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/info", info)
	r.HandleFunc("/", reader)
	r.HandleFunc("/books/{title}/page/{page}", gorilla)

	http.ListenAndServe(":4080", r)
}

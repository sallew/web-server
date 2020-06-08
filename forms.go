// forms.go
package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	Name string
}

type contactDetails struct {
	Email   string
	Subject string
	Message string
}

type student struct {
	id   int
	Name string
}

var details contactDetails
var data user

func names(w http.ResponseWriter, r *http.Request) {
	tmplt := template.Must(template.ParseFiles("names.html"))
	p := student{id: 1, Name: "Aisha"}
	w.Header().Set("Content-Type", "text/html")
	tmplt.Execute(w, p)
}

func hello(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("hello.gohtml"))
	data.Name = "Jon Smith"
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, data)
}

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

	//tmpl.Execute(w, struct{ Success bool }{true})
	http.Redirect(w, r, "/info", http.StatusSeeOther)
	//fmt.Fprintf(w, "_thanks_\n")
}

func gorilla(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/names", names)
	r.HandleFunc("/hello", hello)
	r.HandleFunc("/info", info)
	r.HandleFunc("/", reader)
	r.HandleFunc("/books/{title}/page/{page}", gorilla)

	http.ListenAndServe(":4080", r)
}

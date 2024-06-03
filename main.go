package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	Ascii "ascii-art-wed/banner"
	// Ascii "ascii-art-wed/banner"
)

type Content struct {
	Message string
}

// Template to render the form
var formTemplate = template.Must(template.ParseFiles("form.html"))

func main() {
	// Define HTTP routes
	http.HandleFunc("/ascii", form)

	http.HandleFunc("/", Index)

	// Start the HTTP server
	log.Println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	// http.ServeFile(w, r, "index.html")
	tmpl.Execute(w, nil)
}

func form(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("inputText")
	namebaner := r.FormValue("banner")

	tmpl := template.Must(template.ParseFiles("./form.html"))

	str := strings.Split(name, "\n")
	var data string
	for _, wrd := range str {
		data += Ascii.PrintBanner(wrd, namebaner)
	}

	data = "\n" + data
	content := Content{Message: data}
	tmpl.Execute(w, content)
}

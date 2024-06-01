package main

import (
	"html/template"
	"log"
	"net/http"
)

// Template to render the form
var formTemplate = template.Must(template.ParseFiles("form.html"))

func main() {
	// Define HTTP routes
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/submit", submitHandler)

	// Start the HTTP server
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

// Handler for the form page
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Render the form template
	w.Header().Set("Content-Type", "text/html")
	err := formTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// Handler for form submission
func submitHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get form values
	inputText := r.Form.Get("inputText")
	banner := r.Form.Get("banner")

	// Generate ASCII art (you need to implement this)
	// asciiArt := generateASCIIArt(inputText, banner)

	// Send response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	// Replace the following line with sending the generated ASCII art
	w.Write([]byte("<h1>ASCII Art:</h1><pre>" + inputText + "</pre>"))

}

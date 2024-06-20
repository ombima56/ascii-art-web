package Ascii

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

// ServeError serves a custom error page
func ServeError(w http.ResponseWriter, status int, errMsg string) {
	w.WriteHeader(status)
	tmpl, err := template.ParseFiles("template/error.html")
	if err != nil {
		log.Printf("Error loading error page template: %v", err)
		http.Error(w, "Error loading error page", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, struct{ ErrorMessage string }{ErrorMessage: errMsg})
	if err != nil {
		log.Printf("Error rendering error page template: %v", err)
		http.Error(w, "Error rendering error page", http.StatusInternalServerError)
	}
}

func IndexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ServeError(w, http.StatusNotFound, "404 Page Not Found")
		return
	}
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		ServeError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		ServeError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ServeError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	// Check if the URL path is exactly "/submit"
	if r.URL.Path != "/submit" {
		ServeError(w, http.StatusNotFound, "Page Not Found")
		return
	}

	message := r.FormValue("message")
	bannerfile := r.FormValue("bannerfile")
	if message == "" || bannerfile == "" {
		ServeError(w, http.StatusBadRequest, "Bad Request: Missing message or bannerfile")
		return
	}

	str := strings.Split(message, "\r\n")
	var asciiArt string
	for _, ch := range str {
		asciiArt += PrintBanner(ch, bannerfile)
	}

	tmpl, err := template.ParseFiles("template/result.html")
	if err != nil {
		ServeError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = tmpl.Execute(w, struct{ AsciiArt string }{AsciiArt: asciiArt})
	if err != nil {
		ServeError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

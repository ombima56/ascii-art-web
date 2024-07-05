package Ascii

import (
	"net/http"
	"strings"
	"text/template"
)

func IndexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if the URL path is exactly "/submit"
	if r.URL.Path != "/submit" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	message := r.FormValue("message")
	bannerfile := r.FormValue("bannerfile")
	if message == "" || bannerfile == "" {
		http.Error(w, "400 Bad Request: Missing message or bannerfile", http.StatusBadRequest)
		return
	}

	str := strings.Split(message, "\r\n")
	var asciiArt string
	for _, ch := range str {
		asciiArt += PrintBanner(ch, bannerfile)
	}

	tmpl, err := template.ParseFiles("template/result.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, struct{ AsciiArt string }{AsciiArt: asciiArt})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

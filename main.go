package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	Ascii "ascii-art-wed/asciiArtFunctions"
)

func indexHanleFunc(w http.ResponseWriter, r *http.Request) {
	tml, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err2 := tml.Execute(w, nil)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	message := r.FormValue("message")
	bannerfile := r.FormValue("bannerfile")

	str := strings.Split(message, "\r\n")

	var asciiArt string
	for _, ch := range str {
		asciiArt += Ascii.PrintBanner(ch, bannerfile)
	}

	tml, err := template.ParseFiles("template/result.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err2 := tml.Execute(w, struct{ AsciiArt string }{AsciiArt: asciiArt})
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", indexHanleFunc)
	http.HandleFunc("/submit", submitHandler)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server started at http://localhost:8080")
	err3 := http.ListenAndServe(":8080", nil)
	if err3 != nil {
		log.Fatal(err3)
	}
}

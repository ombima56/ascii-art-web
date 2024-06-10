package main

import (
	Ascii "ascii-art-wed/asciiArtFunctions"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
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

	data := strings.Split(message, "\r\n")

	var asciiArt string
	for _, ch := range data {
		asciiArt += Ascii.PrintBanner(ch, bannerfile)
	}
	fmt.Fprint(w, asciiArt)
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/styles.css")
}

func main() {
	http.HandleFunc("/", indexHanleFunc)
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/styles.css", cssHandler)

	fmt.Println("Server started at http://localhost:8080")
	err3 := http.ListenAndServe(":8080", nil)
	if err3 != nil {
		log.Fatal(err3)
	}
}

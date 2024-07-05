package main

import (
	"log"
	"net/http"

	Ascii "ascii-art-wed/asciiArtFunctions"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Ascii.IndexHandlerFunc)
	mux.HandleFunc("/submit", Ascii.SubmitHandler)
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server Listening on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println("Error starting server", err)
	}
}

package main

import (
	"fmt"
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

	// Handle 404 errors with custom handler
	mux.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		Ascii.ServeError(w, http.StatusNotFound, "Page Not Found")
	})

	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

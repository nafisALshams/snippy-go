package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()
	// Registering handlers
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippy/view/{id}", snippyView)
	mux.HandleFunc("GET /snippy/create", snippyCreate)
	mux.HandleFunc("POST /snippy/create", snippyCreatePost)

	// message Start the server
	log.Print("starting server on :4000")

	// Use http.ListenAndServe to start the server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
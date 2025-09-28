package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Handler for the home page
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go Web Server")
	w.Write([]byte("Welcome to the Home Page!"))
}

// Handler to view a snippy by ID
func snippyView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Viewing a snippy with ID: %d", id)
}

// Handler to create a new snippy
func snippyCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippy here"))
}

// Handler to process the creation of a new snippy
func snippyCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Saving a snippy successfully"))
}

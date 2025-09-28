package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Handler for the home page
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go Web Server")
	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
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

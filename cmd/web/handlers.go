package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// Handler for the home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go Web Server")

	// Use the current working directory as the base for template paths
	cwd, err := os.Getwd()
	if err != nil {
		log.Print("Failed to get working directory:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	files := []string{
		filepath.Join(cwd, "ui", "html", "base.tmpl.html"),
		filepath.Join(cwd, "ui", "html", "partials", "nav.tmpl.html"),
		filepath.Join(cwd, "ui", "html", "pages", "home.tmpl.html"),
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// Handler to view a snippy by ID
func (app *application) snippyView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Viewing a snippy with ID: %d", id)
}

// Handler to create a new snippy
func (app *application) snippyCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippy here"))
}

// Handler to process the creation of a new snippy
func (app *application) snippyCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Saving a snippy successfully"))
}

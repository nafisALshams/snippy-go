package main

import (
	"net/http"
)


func (app *application) routes() *http.ServeMux {
		// Create a new ServeMux
	mux := http.NewServeMux()

	// Serve static files (CSS, JS, etc.)
	fileServer := http.FileServer(http.Dir("ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	// Registering handlers
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippy/view/{id}", app.snippyView)
	mux.HandleFunc("GET /snippy/create", app.snippyCreate)
	mux.HandleFunc("POST /snippy/create", app.snippyCreatePost)
	return mux
}
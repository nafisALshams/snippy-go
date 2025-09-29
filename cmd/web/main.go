package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	//defining a command-line flag for the server address
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Initialize the application struct for dependencies
	app := &application{
		logger: logger,
	}

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

	// message Start the server
	logger.Info("starting server on", "addr", *addr)

	// Use http.ListenAndServe to start the server
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}

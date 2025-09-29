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

	// message Start the server
	logger.Info("starting server on", "addr", *addr)

	// Use http.ListenAndServe to start the server
	// calling app.routes() to get the configured ServeMux from routes.go
	err := http.ListenAndServe(*addr, app.routes()) 
	logger.Error(err.Error())
	os.Exit(1)
}

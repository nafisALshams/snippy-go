package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, This is a simple web server in Go!"))
}

func main() {
	//creating a new server mux and registering a handler function
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	//logging server start
	log.Println("Starting server on :4000")

	//starting the server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

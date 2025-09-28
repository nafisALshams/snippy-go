package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

/* handler function for specific routes start here */

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, This is a simple web server in Go!"))
}

func snippyView(w http.ResponseWriter, r *http.Request) {
	/* Extracting the ID from the URL path */
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("viewing a specific snippy with ID %d", id)
	w.Write([]byte(msg))
}

func snippyCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("creating a snippy"))
}

/* hander function for specific routes end here */

func main() {
	/* creating a new servedxr mux and registering a handler function */
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)                    //restricting to only root path
	mux.HandleFunc("/snippy/view/{id}", snippyView) //dynamic path or wildcard segment {id}
	mux.HandleFunc("/snippy/create", snippyCreate)

	//logging server start
	log.Println("Starting server on :4000")

	//starting the server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

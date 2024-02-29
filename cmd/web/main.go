package main

import (
	"log"
	"net/http"
)

func main() {
	// router
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	//server
	log.Println("Starting server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// create a fileserver which serves file out of "./ui/static" directory
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// use mux.Handle() to register fileServer as a handler
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// route the other application as normal
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	//server
	log.Println("Starting server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}

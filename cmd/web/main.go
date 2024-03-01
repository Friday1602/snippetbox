package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	
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
	log.Printf("Starting server on port %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)

}

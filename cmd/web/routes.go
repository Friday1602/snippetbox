package main

import "net/http"

func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()
	// create a fileserver which serves file out of "./ui/static" directory
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// use mux.Handle() to register fileServer as a handler
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// route the other application as normal
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	return mux
}
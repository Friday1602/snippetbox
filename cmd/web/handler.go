package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// change the signature of func to be a method against *application.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) // use notFound helper instead http.Notfound
		return
	}
	// file path
	files := []string{
		"./ui/html/pages/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}
	// parse template for home
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // use the serverError() helper.
		return
		/*
			app.errorLog.Print(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		*/
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err) // use the serverError() helper.
		/*
			app.errorLog.Print(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		*/
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // use the notFound() helper instead http.Notfound()
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with id %d...", id)

}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed) // use the clientError() helper.
		return
	}

	w.Write([]byte("create a new snippet..."))

}

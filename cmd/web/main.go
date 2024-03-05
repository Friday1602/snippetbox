package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// define application struct to hold the application-wide dependencies.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// create command line flag to receive port
	addr := flag.String("addr", ":4000", "HTTP network address")
	// parse the command line flag
	flag.Parse()
	// create custom logger for info output to standard out(stdout)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// create custom logger for error output to standard error(stderr)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new instance of our application struct, containing
	// the dependencies.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := http.NewServeMux()
	// create a fileserver which serves file out of "./ui/static" directory
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// use mux.Handle() to register fileServer as a handler
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// route the other application as normal
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// create a new http.Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	//use infoLog to log the information
	infoLog.Printf("Starting server on port %s", *addr)
	// Call the ListenAndServe() method instead of calling http.ListenAndServe() func
	err := srv.ListenAndServe()
	//use errorLog to log the error
	errorLog.Fatal(err)

}

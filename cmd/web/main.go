package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/friday1602/snippetbox/internal/models"
	_ "github.com/go-sql-driver/mysql" // need driver's init func so it can register itself with database/sql package
)

// define application struct to hold the application-wide dependencies.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {
	// create command line flag to receive port
	addr := flag.String("addr", ":4000", "HTTP network address")
	// command line flag for MySQL DSN string
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true",
		"MySQL data source name")

	// parse the command line flag
	flag.Parse()
	// create custom logger for info output to standard out(stdout)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// create custom logger for error output to standard error(stderr)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// pass openDB() the dsn fron cl flag
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	// Initialize a new instance of our application struct, containing
	// the dependencies.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

	// create a new http.Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.route(),
	}
	//use infoLog to log the information
	infoLog.Printf("Starting server on port %s", *addr)
	// Call the ListenAndServe() method instead of calling http.ListenAndServe() func
	err = srv.ListenAndServe()
	//use errorLog to log the error
	errorLog.Fatal(err)

}

// openDB() function wraps sql.Open() and return a sql.DB connection pool.
// for given DSN
func openDB(dsn string) (*sql.DB, error) {
	// sql.Open() initialize the pool for future use
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// db.Ping() create connection
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

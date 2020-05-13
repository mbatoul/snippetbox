package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

// Parse the runtime configuration settings for the application
// Establish the dependencies for the handlers
// Run the HTTP server.
func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on port  %s", *addr)
	infoLog.Printf("Listening http://localhost%s", *addr)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}

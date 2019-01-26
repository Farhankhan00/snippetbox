package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	port := flag.String("port", ":4000", "HTTP network port")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.LUTC)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)

	//Use the http.NewServMux function to initialize a new servemux, then
	//register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *port,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just creater. If http.ListenAndServe() return an error
	// we use the log,Fatal() function to log the error message and exit.
	infoLog.Printf("Starting server on %s", *port)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

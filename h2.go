package main

import (
	"flag"
	"log"
	"net/http"
)

// Allow to specify the port via `-port` flag. Default to 8000.
var port = flag.String("port", "8000", "Port")

func main() {
	flag.Parse()
	srv := &http.Server{Addr: ":" + *port, Handler: http.HandlerFunc(handle)}

	// Start the server with TLS, since we are running HTTP/2 it must be
	// run with TLS.
	log.Printf("Serving on https://0.0.0.0:" + *port)
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got connection: %s", r.Proto)

	if r.URL.Path == "/with-content-length" {
		w.Header().Set("Content-Length", "0")
		log.Printf("Returning 'content-length: 0' response header")
	}

	w.WriteHeader(http.StatusNoContent)
}

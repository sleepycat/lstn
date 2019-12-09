package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {

	port := flag.String("port", "3000", "The port to listen on.")
	status := flag.Int("status", 200, "The http status to return.")
	body := flag.String("body", "", "The http response body.")
	path := flag.String("path", "/", "The request path to respond to.")

	flag.Parse()

	http.HandleFunc(*path, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(*status)
		fmt.Fprintf(w, *body)
	})

	http.ListenAndServe(":"+*port, nil)
}

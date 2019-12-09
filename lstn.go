package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {

	port := flag.String("port", "3000", "The port to listen on. Listens on 3000 by default.")
	status := flag.Int("status", 200, "The http status to return. Returns 200 by default.")
	body := flag.String("body", "", "The http response body. Defaults to an empty string.")
	path := flag.String("path", "/", "The request path to respond to. Defaults to /.")

	flag.Parse()

	http.HandleFunc(*path, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(*status)
		fmt.Fprintf(w, *body)
	})

	http.ListenAndServe(":"+*port, nil)
}

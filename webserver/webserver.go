package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a router.
	rt := mux.NewRouter().StrictSlash(true)

	// Add the "index" or root path.
	rt.HandleFunc("/", Index)

	// Fire up the server.
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", rt))
}

// Index This method is the "index" route handler.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from %q", html.EscapeString(r.URL.Path))
}

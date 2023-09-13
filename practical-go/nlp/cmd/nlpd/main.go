package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Routing.
	// /health is an exact match.
	// /health/ is a prefix match.
	http.HandleFunc("/health", healthHandler)

	// Run server.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Run a health check.
	fmt.Fprintln(w, "OK")
}

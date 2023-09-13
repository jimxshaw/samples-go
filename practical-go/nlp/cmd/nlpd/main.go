package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"samples-go/practical-go/nlp"
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

type response struct {
	Tokens []string `json:"tokens"`
}

// tokenizeHandler will read the text from the request body
// and return JSON in the format: "{"tokens": ["hello", "there"]}".
func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	tokens := nlp.Tokenize(string(bodyBytes))

	respTokens := response{
		Tokens: tokens,
	}

	jsonData, err := json.Marshal(respTokens)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Run a health check.
	fmt.Fprintln(w, "OK")
}

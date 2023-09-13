package main

import (
	"encoding/json"
	"expvar"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"samples-go/practical-go/nlp"
	"samples-go/practical-go/nlp/stemmer"
)

var (
	// metrics: http://localhost:8080/debug/vars
	numTok = expvar.NewInt("tokenize.calls")
)

func main() {
	// Create server
	logger := log.New(log.Writer(), "nlp ", log.LstdFlags|log.Lshortfile)
	s := Server{
		logger: logger, // dependency injection
	}

	// Routing.
	// /health is an exact match.
	// /health/ is a prefix match.
	// http.HandleFunc("/health", healthHandler)
	// http.HandleFunc("/tokenize", tokenizeHandler)
	r := mux.NewRouter()
	r.HandleFunc("/health", s.healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/tokenize", s.tokenizeHandler).Methods(http.MethodPost)
	r.HandleFunc("/stem/{word}", s.stemHandler).Methods(http.MethodGet)
	http.Handle("/", r)

	// Run server.
	addr := ":8080"
	s.logger.Printf("server starting on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

type Server struct {
	logger *log.Logger
}

type response struct {
	Tokens []string `json:"tokens"`
}

func (s *Server) stemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]
	stem := stemmer.Stem(word)

	fmt.Fprintln(w, stem)
}

// tokenizeHandler will read the text from the request body
// and return JSON in the format: "{"tokens": ["hello", "there"]}".
func (s *Server) tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	// Before using gorilla/mux:
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	numTok.Add(1)

	// Get, convert and validate the data.
	// In Production, do not just read everything. Add a limit.
	rdr := io.LimitReader(r.Body, 1_000_000)
	bodyBytes, err := io.ReadAll(rdr)
	if err != nil {
		s.logger.Printf("error: can't read - %s", err)
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}

	if len(bodyBytes) == 0 {
		http.Error(w, "missing data", http.StatusBadRequest)
		return
	}

	// Do some work.
	tokens := nlp.Tokenize(string(bodyBytes))

	// Encode and then emit response.
	respTokens := response{
		Tokens: tokens,
	}

	jsonData, err := json.Marshal(respTokens)
	if err != nil {
		http.Error(w, "can't encode", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Run a health check.
	fmt.Fprintln(w, "OK")
}

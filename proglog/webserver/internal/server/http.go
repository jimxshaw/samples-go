package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

/*
A Go web server has a function called HandleFunc(w, *r) for each of the API's endpoints.
This web server has two endpoints, Produce for writing to the commit log and Consume for
reading from the commit log.

A json/http Go server typically has 3 steps:
1) Unmarshal the request's json body into a struct.
2) Run the endpoint's logic with the request to get a result.
3) Marshal and write the result to the response.

More complex web servers will need to move the handlers into http middleware.
*/

type httpServer struct {
	Log *Log
}

type ProduceRequest struct {
	// Record that the API caller wants to
	// append to the log.
	Record Record `json:"record"`
}

type ProduceResponse struct {
	// Tells the caller what offset the
	// log stored the records under.
	Offset uint64 `json:"offset"`
}

type ConsumeRequest struct {
	// States which records the API
	// caller wants to read.
	Offset uint64 `json:"offset"`
}

type ConsumeResponse struct {
	// Sends back the records to the API caller.
	Record Record `json:"record"`
}

func NewHttpServer(address string) *http.Server {
	server := newHttpServer()
	router := mux.NewRouter()
	router.HandleFunc("/", server.handleProduce).Methods("POST")
	router.HandleFunc("/", server.handleConsume).Methods("GET")

	return &http.Server{
		Addr:    address,
		Handler: router,
	}
}

func newHttpServer() *httpServer {
	return &httpServer{
		Log: NewLog(),
	}
}

func (s *httpServer) handleProduce(w http.ResponseWriter, r *http.Request) {
	var request ProduceRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	offset, err := s.Log.Append(request.Record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ProduceResponse{Offset: offset}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *httpServer) handleConsume(w http.ResponseWriter, r *http.Request) {
	var request ConsumeRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	record, err := s.Log.Read(request.Offset)
	if err == ErrOffsetNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ConsumeResponse{Record: record}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

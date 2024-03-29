package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// PageData - This is a basic struct to hold page
// data variables.
type PageData struct {
	Title string
	Body  string
}

// Define a basic html template.
const html = `
			<html>
			<head><title>{{.Title}}</title></head>
			<body>
			<h1>{{.Title}}</h1>
			{{.Body}}
			</body>
			</html>
			`

func main() {
	// Create a router.
	rt := mux.NewRouter().StrictSlash(true)

	// Add the "index" or root path.
	rt.HandleFunc("/", Index)

	// Fire up the server.
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", rt))
}

// Index - This method is the "index" route handler.
func Index(w http.ResponseWriter, r *http.Request) {
	// Fill out page data for index.
	pd := PageData{
		Title: "Index Page",
		Body:  "This is the body of the page.",
	}

	// Render a template with our page data.
	tmpl, err := render(pd)

	// If we get an error, write it out and exit.
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// If all went well, write out the template.
	w.Write([]byte(tmpl))
}

func render(pd PageData) (string, error) {
	// Parse the template.
	tmpl, err := template.New("index").Parse(html)

	if err != nil {
		return "", err
	}

	// We need somewhere to write out the executed template.
	var out bytes.Buffer

	// Render the template with the data we passed in.
	if err := tmpl.Execute(&out, pd); err != nil {
		// If we couldn't render, return an error.
		return "", err
	}

	// Return the template.
	return out.String(), nil
}

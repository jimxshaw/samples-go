package main

import (
	"log"

	"github.com/jimxshaw/samples-go/proglog/webserver/internal/server"
)

func main() {
	server := server.NewHttpServer(":8000")
	log.Fatal(server.ListenAndServe())
}

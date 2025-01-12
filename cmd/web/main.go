package main

import (
	"go-websockets/internal/handlers"
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener.")
	go handlers.ReadFromWsChannel()

	log.Println("Starting the server on port 8080")

	must(http.ListenAndServe(":8080", mux))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

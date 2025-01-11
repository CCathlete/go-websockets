package main

import (
	"go-websockets/internal/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func routes() http.Handler {
	// We're using a pattern mux to have flexibility in our urls.
	mux := pat.New()

	mux.Get("/", handlers.Home)
	mux.Get("/ws", handlers.WsEndpoint)

	return mux
}

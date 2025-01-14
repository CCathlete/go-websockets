package main

import (
	"go-websockets/internal/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func routes() http.Handler {
	// We're using a pattern mux to have flexibility in our urls.
	mux := pat.New()

	// Setting up a file server object to serve static js files.
	// It has an inner filesystem with a root folder at ./html/typescript.
	staticJsFileHandler := http.FileServer(http.Dir("./html/typescript"))

	// Adjusting the request url to look like a path inside ./html/typescript.
	staticJsFileHandler =
		http.StripPrefix("/typescript/", staticJsFileHandler)

	// Our html file, located in ./html will send a get request to /typescript/<file>.js, the prefix will get stripped and the handler will look for the file inside its root folder.
	mux.Get("/typescript/", staticJsFileHandler)
	mux.Get("/", handlers.Home)
	mux.Get("/ws", handlers.WsEndpoint)

	return mux
}

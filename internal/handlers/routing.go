package handlers

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func routes() http.Handler {
	// We're using a pattern mux to have flexibility in our urls.
	mux := pat.New()

	// Converting a function with the appropriate signatore to a handlerFunc type which is a funciton that satisfies the http.Handler interface.
	mux.Get("/", http.HandlerFunc(Home))

	return mux
}

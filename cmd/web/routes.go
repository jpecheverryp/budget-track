package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", app.handler.Index)
	mux.HandleFunc("GET /login", app.handler.Login)

	return mux
}

package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.getIndex)
	mux.HandleFunc("GET /test", app.getTest)
	mux.HandleFunc("GET /dashboard", app.getDashboard)
	mux.HandleFunc("GET /transactions/view/{id}", app.getTransactionView)
	mux.HandleFunc("GET /transactions/create", app.getTransactionCreate)

	mux.HandleFunc("POST /accounts/create", app.postAccountCreate)

	mux.HandleFunc("GET /auth/register", app.getRegister)
	mux.HandleFunc("POST /auth/register", app.postRegister)

	return commonHeaders(mux)
}

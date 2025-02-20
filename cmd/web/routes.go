package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.getIndex)
	mux.HandleFunc("GET /dashboard", app.getDashboard)
	mux.HandleFunc("GET /transactions/view/{id}", app.getTransactionView)
	mux.HandleFunc("GET /transactions/create", app.getTransactionCreate)
	mux.HandleFunc("POST /transactions/create", app.postTransactionCreate)

	mux.HandleFunc("GET /auth/register", app.getRegister)
    mux.HandleFunc("POST /auth/register", app.postRegister)

	return mux
}

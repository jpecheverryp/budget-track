package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	//File server will go here, after adding static

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.getIndex))
	mux.Handle("GET /test", dynamic.ThenFunc(app.getTest))
	mux.Handle("GET /dashboard", dynamic.ThenFunc(app.getDashboard))

	mux.Handle("POST /accounts/create", dynamic.ThenFunc(app.postAccountCreate))

	mux.Handle("GET /auth/register", dynamic.ThenFunc(app.getRegister))
	mux.Handle("POST /auth/register", dynamic.ThenFunc(app.postRegister))
	mux.Handle("GET /auth/login", dynamic.ThenFunc(app.getLogin))
	mux.Handle("POST /auth/login", dynamic.ThenFunc(app.postLogin))
	mux.Handle("POST /auth/logout", dynamic.ThenFunc(app.postLogout))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	return standard.Then(mux)
}

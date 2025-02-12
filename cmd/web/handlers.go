package main

import (
	"net/http"

	"budget-track.jpech.dev/views"
)

func showIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func showLogin(w http.ResponseWriter, r *http.Request) {
	component := views.Login()
	component.Render(r.Context(), w)
}

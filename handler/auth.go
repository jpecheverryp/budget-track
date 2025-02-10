package handler

import (
	"net/http"

	"budget-track.jpech.dev/views"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	component := views.Login()
	component.Render(r.Context(), w)
}

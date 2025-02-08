package main

import (
	"net/http"

	"budget-track.jpech.dev/views/layout"
)

func (app *application) routes() http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        component := layout.Base()
        component.Render(r.Context(), w)
    })

    return mux
}

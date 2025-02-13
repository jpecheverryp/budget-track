package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home.html", templateData{})
}

func (app *application) getTransactionView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Showing a single transaction info with id: %d", id)
}

func (app *application) getTransactionCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show page to add transaction"))
}
func (app *application) postTransactionCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new transaction"))
}

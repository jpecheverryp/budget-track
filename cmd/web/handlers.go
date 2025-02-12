package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
        app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
        app.serverError(w, r, err)
		return
	}
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

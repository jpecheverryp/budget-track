package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"budget-track.jpech.dev/store"
	"github.com/google/uuid"
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
	accountId, err := uuid.NewV6()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = app.transactions.Insert(&store.Transaction{
		Description:     "laptop",
		AccountID:       accountId,
		ValueInCents:    100000,
		TransactionDate: time.Now(),
	})

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.logger.Info("created new transaction")
	w.Write([]byte("Save a new transaction"))

}

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"budget-track.jpech.dev/views/page"
)

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home.html", templateData{})
}

func (app *application) getTest(w http.ResponseWriter, r *http.Request) {
	accounts, err := app.repo.ReadAllAccounts(r.Context())
	if err != nil {
		app.serverError(w, r, err)
	}
	app.render(w, r, http.StatusOK, "test.html", templateData{
		Accounts: accounts,
	})
}

func (app *application) getDashboard(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "dashboard.html", templateData{})
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

func (app *application) postAccountCreate(w http.ResponseWriter, r *http.Request) {
	_, err := app.repo.CreateAccount(r.Context(), "Chase")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.logger.Info("created new transaction")
	http.Redirect(w, r, "/test", http.StatusSeeOther)

}

type userRegisterForm struct {
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (app *application) getRegister(w http.ResponseWriter, r *http.Request) {
	err := page.Register().Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) postRegister(w http.ResponseWriter, r *http.Request) {
	var form userRegisterForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

}

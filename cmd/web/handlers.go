package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jpecheverryp/budget-track/internal/service"
	"github.com/jpecheverryp/budget-track/views/page"
)

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home.html", templateData{})
}

func (app *application) getTest(w http.ResponseWriter, r *http.Request) {
	accounts, err := app.accounts.GetAll(r.Context())

	if err != nil {
		app.serverError(w, r, err)
		return
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
	accountInput := service.AccountCreateInput{
		AccountName: "Fidelity",
	}
	_, err := app.accounts.Insert(r.Context(), accountInput)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.logger.Info("created new transaction")
	http.Redirect(w, r, "/test", http.StatusSeeOther)

}

type userRegisterForm struct {
	Username string
	Email    string
	Password string
}

func (app *application) getRegister(w http.ResponseWriter, r *http.Request) {
	err := page.Register().Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) postRegister(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	registerFormData := &userRegisterForm{
		Username: r.Form.Get("username"),
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	app.logger.Info("form: ", "username", registerFormData.Username)
	app.logger.Info("form: ", "email", registerFormData.Email)
	app.logger.Info("form: ", "password", registerFormData.Password)

	app.logger.Info("user created succesfully")

	http.Redirect(w, r, "/test", http.StatusSeeOther)
}

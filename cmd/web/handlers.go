package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/jpecheverryp/budget-track/internal/service"
	"github.com/jpecheverryp/budget-track/views/page"
)

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	component := page.Home()
	err := component.Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getTest(w http.ResponseWriter, r *http.Request) {
	accounts, err := app.service.Accounts.GetAll(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	flash := app.sessionManager.PopString(r.Context(), "flash")

	pageData := page.TestPageData{
		Accounts: accounts,
		Flash:    flash,
	}

	component := page.TestAPI(pageData)
	err = component.Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getDashboard(w http.ResponseWriter, r *http.Request) {
	component := page.MainDashboard()
	err := component.Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
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

func (app *application) postAccountCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	accountInput := service.AccountCreateInput{
		AccountName: r.Form.Get("account-name"),
	}

	_, err = app.service.Accounts.Insert(r.Context(), accountInput)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.logger.Info("created new account")
	app.sessionManager.Put(r.Context(), "flash", "account added succesfully!")
	http.Redirect(w, r, "/test", http.StatusSeeOther)

}

type registerFormData struct {
	Values struct {
		Username string
		Email    string
		Password string
	}
	Errors struct {
		Username string
		Email    string
		Password string
	}
}

func (f *registerFormData) Validate() bool {
	var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	isValidForm := true

	// Check Username
	if f.Values.Username == "" {
		f.Errors.Username = "This field cannot be blank"
		isValidForm = false
	}

	// Check Email
	if f.Values.Email == "" {
		f.Errors.Email = "This field cannot be blank"
		isValidForm = false
	}
	if !EmailRX.MatchString(f.Values.Email) {
		f.Errors.Email = "This field must be a valid email address"
		isValidForm = false
	}

	// Check Password
	if f.Values.Password == "" {
		f.Errors.Password = "This field cannot be blank"
		isValidForm = false
	}

	return isValidForm
}

// Show form to register
func (app *application) getRegister(w http.ResponseWriter, r *http.Request) {
	err := page.Register().Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// Create new user
func (app *application) postRegister(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	registerFormData := &registerFormData{
		Values: struct {
			Username string
			Email    string
			Password string
		}{
			Username: r.Form.Get("username"),
			Email:    r.Form.Get("email"),
			Password: r.Form.Get("password"),
		},
	}

	isValid := registerFormData.Validate()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	if !isValid {
		component := page.Register()
		err := component.Render(r.Context(), w)
		if err != nil {
			app.serverError(w, r, err)
		}
		return
	}

	app.logger.Info("user created succesfully")
	http.Redirect(w, r, "/test", http.StatusSeeOther)
}

type loginFormData struct {
	Username string
	Email    string
	Password string
}

func (app *application) getLogin(w http.ResponseWriter, r *http.Request) {
	err := page.Login().Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// Create new user
func (app *application) postLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	loginFormData := &loginFormData{
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	app.logger.Info("form: ", "email", loginFormData.Email)
	app.logger.Info("form: ", "password", loginFormData.Password)

	app.logger.Info("user authenticated succesfully")

	http.Redirect(w, r, "/test", http.StatusSeeOther)
}

func (app *application) postLogout(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("user logged out succesfully")
}

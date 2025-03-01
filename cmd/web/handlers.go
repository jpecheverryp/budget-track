package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jpecheverryp/budget-track/internal/repository"
	"github.com/jpecheverryp/budget-track/internal/validator"
	"github.com/jpecheverryp/budget-track/views/page"
	"golang.org/x/crypto/bcrypt"
)

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	component := page.Home()
	err := component.Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getTest(w http.ResponseWriter, r *http.Request) {
	accounts, err := app.repo.ReadAllAccounts(r.Context())
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

	accountName := r.Form.Get("account-name")

	_, err = app.repo.CreateAccount(r.Context(), accountName)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.logger.Info("created new account")
	app.sessionManager.Put(r.Context(), "flash", "account added succesfully!")
	http.Redirect(w, r, "/test", http.StatusSeeOther)

}

// Show form to register
func (app *application) getRegister(w http.ResponseWriter, r *http.Request) {
	registerFormData := page.RegisterFormData{}
	err := page.Register(registerFormData).Render(r.Context(), w)
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

	form := page.RegisterFormData{
		Username: r.Form.Get("username"),
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	form.CheckField(validator.NotBlank(form.Username), "username", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Email), "username", "This field cannot be blank")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
	form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
	form.CheckField(validator.MinChars(form.Password, 8), "password", "This field must be at least 8 characters long")

	emailCount, err := app.repo.EmailTaken(r.Context(), form.Email)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	if emailCount > 0 {
		form.AddFieldError("email", "Email already taken")
	}

	if !form.Valid() {
		component := page.Register(form)
		err := component.Render(r.Context(), w)
		if err != nil {
			app.serverError(w, r, err)
		}
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), 12)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	newUser := repository.RegisterUserParams{
		Username:     form.Username,
		Email:        form.Email,
		PasswordHash: string(hashedPassword),
	}
	err = app.repo.RegisterUser(r.Context(), newUser)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.logger.Info("user created succesfully")
	http.Redirect(w, r, "/test", http.StatusSeeOther)
}

func (app *application) getLogin(w http.ResponseWriter, r *http.Request) {
	loginFormData := page.LoginFormData{}
	err := page.Login(loginFormData).Render(r.Context(), w)
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

	form := page.LoginFormData{
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	authData, err := app.repo.GetAuthByEmail(r.Context(), form.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			form.Validator.AddNonFieldError("InvalidCredentials")
			err := page.Login(form).Render(r.Context(), w)
			if err != nil {
				app.serverError(w, r, err)
			}
			return
		}
		app.serverError(w, r, err)
		return
	}

	app.logger.Info("form: ", "email", form.Email)
	app.logger.Info("form: ", "password", form.Password)
	app.logger.Info("form: ", "hash", authData.PasswordHash)

	app.logger.Info("user authenticated succesfully")

	http.Redirect(w, r, "/test", http.StatusSeeOther)
}

func (app *application) postLogout(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("user logged out succesfully")
}

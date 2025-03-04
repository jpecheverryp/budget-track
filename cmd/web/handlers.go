package main

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/jpecheverryp/budget-track/internal/repository"
	"github.com/jpecheverryp/budget-track/internal/validator"
	"github.com/jpecheverryp/budget-track/views"
	"github.com/jpecheverryp/budget-track/views/page"
	"golang.org/x/crypto/bcrypt"
)

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	pageData := views.PageData{}
	component := page.Home(pageData)
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
	userAccounts, err := app.repo.GetAllUsers(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	flash := app.sessionManager.PopString(r.Context(), "flash")

	pageData := views.PageData{
		Accounts:     accounts,
		UserAccounts: userAccounts,
		Flash:        flash,
	}

	component := page.TestAPI(pageData)
	err = component.Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getDashboard(w http.ResponseWriter, r *http.Request) {
	pageData := views.PageData{}
	component := page.MainDashboard(pageData)
	err := component.Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
	}
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
	pageData := views.PageData{}
	err := page.Register(registerFormData, pageData).Render(r.Context(), w)
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
		pageData := views.PageData{}
		component := page.Register(form, pageData)
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
	pageData := views.PageData{}
	err := page.Login(loginFormData, pageData).Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// Login an existing user
func (app *application) postLogin(w http.ResponseWriter, r *http.Request) {
	// Parse user credentials
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Store credentials
	form := page.LoginFormData{
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	// Validate credentials are right format
	form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
	form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")

	if !form.Valid() {
		pageData := views.PageData{}
		component := page.Login(form, pageData)
		err := component.Render(r.Context(), w)
		if err != nil {
			app.serverError(w, r, err)
		}
		return
	}

	// Get ID and Password by email
	authData, err := app.repo.GetAuthByEmail(r.Context(), form.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No user found, render login
			form.Validator.AddNonFieldError("InvalidCredentials")
            pageData := views.PageData{}
			err := page.Login(form, pageData).Render(r.Context(), w)
			if err != nil {
				app.serverError(w, r, err)
			}
			return
		}
		app.serverError(w, r, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(authData.PasswordHash), []byte(form.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			form.AddNonFieldError("Invalid credentials")
            pageData := views.PageData{}
			err := page.Login(form, pageData).Render(r.Context(), w)
			if err != nil {
				app.serverError(w, r, err)
			}
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Put(r.Context(), "authenticatedUserID", authData.ID.String())

	app.logger.Info("user authenticated succesfully")

	http.Redirect(w, r, "/test", http.StatusSeeOther)
}

func (app *application) postLogout(w http.ResponseWriter, r *http.Request) {
	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Remove(r.Context(), "authenticatedUserID")
	app.sessionManager.Put(r.Context(), "flash", "You've been logged out succesfully!")
	app.logger.Info("user logged out succesfully")

	http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
}

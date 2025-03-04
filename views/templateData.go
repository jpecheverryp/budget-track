package views

import "github.com/jpecheverryp/budget-track/internal/repository"

type PageData struct {
	Flash        string
	Accounts     []repository.Account
	UserAccounts []repository.GetAllUsersRow
}

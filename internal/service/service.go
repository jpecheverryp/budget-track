package service

import "github.com/jpecheverryp/budget-track/internal/repository"

type Service struct {
	Accounts AccountService
	Users    UserService
}

func New(repo repository.Queries) Service {
	return Service{
		Accounts: AccountService{repo: repo},
		Users:    UserService{repo: repo},
	}
}

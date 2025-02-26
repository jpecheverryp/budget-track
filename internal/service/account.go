package service

import (
	"context"

	"github.com/jpecheverryp/budget-track/internal/repository"
)

type AccountService struct {
	repo repository.Queries
}

func New(repo repository.Queries) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (a *AccountService) GetAll(ctx context.Context) ([]repository.Account, error) {
	accounts, err := a.repo.ReadAllAccounts(ctx)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (a *AccountService) Insert(ctx context.Context) (repository.Account, error) {
	account, err := a.repo.CreateAccount(ctx, "Wealthfront")
	if err != nil {
		return repository.Account{}, err
	}

	return account, nil
}

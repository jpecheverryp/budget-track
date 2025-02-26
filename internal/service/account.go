package service

import (
	"context"

	"github.com/jpecheverryp/budget-track/internal/repository"
)

type AccountService struct {
	repo repository.Queries
}

func (a *AccountService) GetAll(ctx context.Context) ([]repository.Account, error) {
	accounts, err := a.repo.ReadAllAccounts(ctx)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

type AccountCreateInput struct {
	AccountName string
}

func (a *AccountService) Insert(ctx context.Context, accountData AccountCreateInput) (repository.Account, error) {
	account, err := a.repo.CreateAccount(ctx, accountData.AccountName)
	if err != nil {
		return repository.Account{}, err
	}

	return account, nil
}

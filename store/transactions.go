package store

import (
	"database/sql"

	"github.com/google/uuid"
)

type TransactionModelInterface interface {
	Insert(description string, valueInCents int) (Transaction, error)
}

type Transaction struct {
	ID           uuid.UUID
	Description  string
	ValueInCents int
}

type TransactionModel struct {
	DB *sql.DB
}

func (m *TransactionModel) Insert(description string, valueInCents int) (Transaction, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return Transaction{}, err
	}
	return Transaction{
		ID:           id,
		Description:  description,
		ValueInCents: valueInCents,
	}, nil
}

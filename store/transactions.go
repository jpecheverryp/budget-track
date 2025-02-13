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

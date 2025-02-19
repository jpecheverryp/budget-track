package store

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type TransactionModelInterface interface {
	Insert(*Transaction) error
}

type Transaction struct {
	ID              uuid.UUID
	Description     string
	AccountID       uuid.UUID
	ValueInCents    int
	TransactionDate time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type TransactionModel struct {
	DB *sql.DB
}

func (m *TransactionModel) Insert(transaction *Transaction) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	query := `
    INSERT INTO transaction (id, description, account_id, value_in_cents, transaction_date)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING (created_at)
    `
	args := []any{id, transaction.Description, transaction.AccountID, transaction.ValueInCents, transaction.TransactionDate}

	return m.DB.QueryRow(query, args...).Scan(&transaction.CreatedAt)
}

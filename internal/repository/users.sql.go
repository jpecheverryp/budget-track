// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const emailTaken = `-- name: EmailTaken :one
SELECT COUNT(*) FROM user_account WHERE email = $1
`

func (q *Queries) EmailTaken(ctx context.Context, email string) (int64, error) {
	row := q.db.QueryRowContext(ctx, emailTaken, email)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getAuthByEmail = `-- name: GetAuthByEmail :one
SELECT id, password_hash FROM user_account WHERE email = $1
`

type GetAuthByEmailRow struct {
	ID           uuid.UUID
	PasswordHash string
}

func (q *Queries) GetAuthByEmail(ctx context.Context, email string) (GetAuthByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getAuthByEmail, email)
	var i GetAuthByEmailRow
	err := row.Scan(&i.ID, &i.PasswordHash)
	return i, err
}

const registerUser = `-- name: RegisterUser :exec
INSERT INTO user_account (id, username, email, password_hash)
VALUES (uuid_generate_v4(), $1, $2, $3)
`

type RegisterUserParams struct {
	Username     string
	Email        string
	PasswordHash string
}

func (q *Queries) RegisterUser(ctx context.Context, arg RegisterUserParams) error {
	_, err := q.db.ExecContext(ctx, registerUser, arg.Username, arg.Email, arg.PasswordHash)
	return err
}

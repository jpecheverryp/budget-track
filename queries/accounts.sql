-- name: CreateAccount :one
INSERT INTO account (id, name, user_id)
VALUES (uuid_generate_v4(), $1, uuid_generate_v4())
RETURNING *;

-- name: ReadAllAccounts :many
SELECT * FROM account;

-- name: RegisterUser :exec
INSERT INTO user_account (id, username, email, password_hash)
VALUES (uuid_generate_v4(), $1, $2, $3);

-- name: EmailTaken :one
SELECT COUNT(*) FROM user_account WHERE email = $1;

-- name: GetAuthByEmail :one
SELECT id, password_hash FROM user_account WHERE email = $1;

-- name: GetAllUsers :many
SELECT id, username, email FROM user_account;

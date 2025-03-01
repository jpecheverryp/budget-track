-- name: TransactionInsert :exec

INSERT INTO transaction (id, description, account_id, value_in_cents, transaction_date)
VALUES (generate_uuid_v4, $1, $2, $3, $4);

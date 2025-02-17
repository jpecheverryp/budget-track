CREATE TABLE IF NOT EXISTS account (
    id uuid PRIMARY KEY,
    name text NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_account (
    id uuid PRIMARY KEY,
    username text NOT NULL,
    email text NOT NULL,
    password_hash text NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS transaction (
    id uuid PRIMARY KEY,
    description text NOT NULL,
    account_id uuid NOT NULL,
    value_in_cents integer NOT NULL,
    transaction_date date NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
);

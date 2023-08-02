-- +goose Up
ALTER TABLE users DROP COLUMN password;

ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
    encode(sha256(random()::text::bytea), 'hex')
);

-- +goose Down
ALTER TABLE DROP COLUMN api_key;
ALTER TABLE ADD COLUMN password TEXT;

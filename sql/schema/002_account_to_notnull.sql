-- +goose Up

ALTER TABLE entries
ALTER COLUMN account_id SET NOT NULL;

-- +goose Down

ALTER TABLE entries
ALTER COLUMN account_id DROP NOT NULL;
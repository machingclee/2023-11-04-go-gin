-- +goose Up

ALTER TABLE accounts DROP COLUMN _deprecated_owner;
ALTER TABLE accounts ALTER COLUMN owner SET NOT NULL;

-- +goose Down


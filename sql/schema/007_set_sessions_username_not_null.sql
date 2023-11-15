-- +goose Up

ALTER TABLE "sessions"
ALTER COLUMN username SET NOT NULL;

-- +goose Down

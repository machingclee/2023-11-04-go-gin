-- +goose Up

ALTER TABLE "accounts" RENAME COLUMN "owner" to "_deprecated_owner";
ALTER TABLE "accounts" ALTER COLUMN "_deprecated_owner" DROP NOT NULL;
ALTER TABLE accounts ADD "owner" varchar;
CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");
CREATE UNIQUE INDEX owner_currency ON "accounts" ("owner", "currency");

-- +goose Down

ALTER TABLE IF EXISTS "accounts" DROP CONTRAINT IF EXISTS "owner_currency";
ALTER TABLE IF EXISTS "accounts" DROP CONTRAINT IF EXISTS "accounts_owner_fkey";
DROP TABLE IF EXISTS "users";
ALTER TABLE "accounts" DROP COLUMN owner;
ALTER TABLE "accounts" ALTER COLUMN "_deprecated_owner" SET NOT NULL;
ALTER TABLE "accounts" RENAME COLUMN "_deprecated_owner" to "owner";
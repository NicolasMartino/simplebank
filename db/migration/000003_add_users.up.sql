CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "password_change_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" 
RENAME COLUMN "owner" TO "user_id";

ALTER TABLE "accounts"
ALTER COLUMN "user_id" TYPE bigint USING user_id::bigint;

ALTER TABLE "accounts" ADD CONSTRAINT "accounts_user_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id");

--CREATE UNIQUE INDEX ON "accounts" ("owner", "currency"); --composite index => one account per currency
ALTER TABLE "accounts" ADD CONSTRAINT "user_currency_key" UNIQUE ("user_id", "currency")
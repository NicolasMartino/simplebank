ALTER TABLE "transfers"
ALTER COLUMN "amount" TYPE bigint using "amount"::integer;

ALTER TABLE "entries"
ALTER COLUMN "amount" TYPE bigint using "amount"::integer;
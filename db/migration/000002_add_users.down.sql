DROP INDEX IF EXISTS "transfers_from_account_id_to_account_id_idx";

DROP INDEX IF EXISTS "transfers_to_account_id_idx";

DROP INDEX IF EXISTS "transfers_from_account_id_idx";

DROP INDEX IF EXISTS "entries_account_id_idx";

DROP INDEX IF EXISTS "accounts_owner_idx";

ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "owner_currency_key";

ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

DROP TABLE IF EXISTS "users";

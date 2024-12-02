ALTER TABLE IF EXISTS "emails" DROP CONSTRAINT IF EXISTS "sessions_user_id_fkey";

DROP TABLE IF EXISTS "emails";
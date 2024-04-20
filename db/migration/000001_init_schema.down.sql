ALTER TABLE IF EXISTS "sessions" DROP CONSTRAINT IF EXISTS "sessions_user_id_fkey";

ALTER TABLE IF EXISTS "recruiters" DROP CONSTRAINT IF EXISTS "sessions_user_id_fkey";

DROP TABLE IF EXISTS "sessions";
DROP TABLE IF EXISTS "recruiters";
DROP TABLE IF EXISTS "users";
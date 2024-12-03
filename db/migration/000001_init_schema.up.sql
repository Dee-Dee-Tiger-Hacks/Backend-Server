CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL UNIQUE,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone" varchar NOT NULL,
  "job_title" varchar NOT NULL,
  "avatar_url" varchar NOT NULL,
  "linkedin_url" varchar NOT NULL,
  "is_email_verified" boolean NOT NULL DEFAULT false,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "create_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT 'now()'
);


CREATE TABLE "recruiters" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "linkedin_url" varchar NOT NULL,
  "name" varchar NOT NULL,
  "company" varchar NOT NULL,
  "email" varchar NOT NULL,
  "overview" varchar NOT NULL,
  "suggested_email" varchar NOT NULL,
  "potential_topics" varchar NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "verify_emails" (
  "id" bigserial PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "sessions" ("user_id", "create_at");

CREATE INDEX ON "recruiters" ("user_id");

ALTER TABLE "verify_emails" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "recruiters" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
CREATE TABLE "emails" (
  "id" uuid NOT NULL PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "subject" varchar NOT NULL,
  "content" varchar NOT NULL,
  "title" varchar NOT NULL,
  "email_address" varchar NOT NULL,
  "create_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "emails" ("user_id");

ALTER TABLE "emails" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE TABLE "resumes" (
  "id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "resume_public_id" varchar NOT NULL,
  "resume_title" varchar NOT NULL,
  "resume_pdf_url" varchar NOT NULL,
   "create_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "resumes" ("user_id");

ALTER TABLE "resumes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

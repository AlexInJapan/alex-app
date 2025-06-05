CREATE TYPE "task_completation" AS ENUM (
  'notStarted',
  'Working',
  'Done'
);

CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL
);

CREATE TABLE "tasks" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "status" task_completation,
  "user_id" bigint
);

CREATE INDEX ON "users" ("name");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "tasks" ("name");

CREATE INDEX ON "tasks" ("status");

CREATE INDEX ON "tasks" ("user_id");

ALTER TABLE "tasks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

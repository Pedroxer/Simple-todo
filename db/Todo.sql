CREATE TABLE "tasks" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "description" text,
  "important" integer,
  "done" integer,
  "deadline" timestamp NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "lists" (
  "id" bigserial PRIMARY KEY,
  "title" varchar UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "user_to_list" (
  "list_id" integer NOT NULL,
  "user_id" integer NOT NULL
);

CREATE TABLE "tasks_to_list" (
  "task_id" integer NOT NULL,
  "list_id" integer NOT NULL
);

COMMENT ON COLUMN "tasks"."important" IS '1 for important, 0 for regular task';

COMMENT ON COLUMN "tasks"."done" IS '1 for done, 0 for in progress ';

ALTER TABLE "user_to_list" ADD FOREIGN KEY ("list_id") REFERENCES "lists" ("id");

ALTER TABLE "user_to_list" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "tasks_to_list" ADD FOREIGN KEY ("list_id") REFERENCES "lists" ("id");

ALTER TABLE "tasks_to_list" ADD FOREIGN KEY ("task_id") REFERENCES "tasks" ("id");

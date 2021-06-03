CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar(32) NOT NULL,
  "last_name" varchar(32) NOT NULL,
  "user_name" varchar(32) NOT NULL,
  "email" varchar(32) NOT NULL,
  "password" varchar(255) NOT NULL,
  "phone_no" integer NOT NULL,
  "ts" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("first_name");
CREATE INDEX ON "users" ("user_name");
CREATE INDEX ON "users" ("email");
CREATE INDEX ON "users" ("phone_no");

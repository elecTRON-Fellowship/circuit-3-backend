CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" text NOT NULL,
  "last_name" text NOT NULL,
  "user_name" text UNIQUE NOT NULL,
  "email" text UNIQUE NOT NULL,
  "password" text NOT NULL,
  "phone_no" text UNIQUE NOT NULL,
  "created_at" timestamptz DEFAULT (now())
  -- "wallet_id" text,
  -- "wallet_reference_id" text
);

-- CREATE TABLE "address" (
--   "id" text PRIMARY KEY,
--   "user_id" bigint,
--   "name" text NOT NULL,
--   "line_1" text NOT NULL,
--   "line_2" text,
--   "line_3" text,
--   "city" text NOT NULL,
--   "district" text,
--   "state" text NOT NULL,
--   "country" text NOT NULL,
--   "zip" text NOT NULL,
--   "phone_no" text NOT NULL,
--   "metadata" jsonb,
--   "created_at" int NOT NULL,
--   "updated_at" int NOT NULL
-- );

-- ALTER TABLE "address" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

-- COMMENT ON COLUMN "users"."phone_no" IS 'Should be in E.164 format';

-- COMMENT ON COLUMN "address"."id" IS 'Should start with address_';

-- COMMENT ON COLUMN "address"."phone_no" IS 'Should be in E.164 format';

-- CREATE INDEX ON "users" ("first_name");
-- CREATE INDEX ON "users" ("user_name");
-- CREATE INDEX ON "users" ("email");
-- CREATE INDEX ON "users" ("phone_no");
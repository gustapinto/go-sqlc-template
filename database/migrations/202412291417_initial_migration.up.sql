BEGIN;

CREATE TABLE IF NOT EXISTS "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "login" VARCHAR(100) UNIQUE NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "email" VARCHAR(254),
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "user_logs" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGINT REFERENCES "users" ("id") NOT NULL,
    "type" VARCHAR(100) NOT NULL,
    "message" VARCHAR(255),
    "created_at" TIMESTAMP NOT NULL
);

COMMIT;

-- name: SelectUserByID :one
SELECT
    id,
    login,
    password,
    email,
    created_at,
    updated_at
FROM
    "users"
WHERE
    id = @id::BIGINT;

-- name: SelectUserByLogin :one
SELECT
    id,
    login,
    password,
    email,
    created_at,
    updated_at
FROM
    "users"
WHERE
    login = @login::VARCHAR;

-- name: InsertUser :one
INSERT INTO "users" (
    login,
    password,
    email,
    created_at
) VALUES (
    @login::VARCHAR,
    @password::VARCHAR,
    @email::VARCHAR,
    @created_at::TIMESTAMP
)
RETURNING *;

-- name: UpdateUserByID :exec
UPDATE
    "users"
SET
    login = @login::VARCHAR,
    email = @email::VARCHAR,
    updated_at = @updated_at::TIMESTAMP
WHERE
    id = @id::BIGINT;


-- name: InsertUserLog :one
INSERT INTO "user_logs" (
    "user_id",
    "type",
    "message",
    "created_at"
)
VALUES (
    @user_id::BIGINT,
    @type::VARCHAR,
    @message::VARCHAR,
    @created_at::TIMESTAMP
)
RETURNING *;

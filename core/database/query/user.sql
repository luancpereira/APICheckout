
-----------------
---- INSERTS ----
-----------------

-- name: InsertUser :one
INSERT INTO "user" (
    email,
    name,
    password,
    permission,
    created_at,
    token_confirmation,
    token_confirmation_expiration_date
) VALUES (
    @email::VARCHAR,
    @name::VARCHAR,
    @password::VARCHAR,
    @permission::VARCHAR,
    @created_at::TIMESTAMP,
    @token_confirmation::VARCHAR,
    @token_confirmation_expiration_date::TIMESTAMP
)
RETURNING id, email;

-----------------
---- INSERTS ----
-----------------


-----------------
---- SELECTS ----
-----------------

-- name: SelectUserForLogin :one
SELECT
    id,
    email,
    "password"::VARCHAR
FROM
    "user"
WHERE
    LOWER(email) = LOWER(@email::VARCHAR);

-- name: SelectUserIDByEmail :one
SELECT
    id::BIGINT
FROM
    "user"
WHERE
    email = @email::VARCHAR;

-----------------
---- SELECTS ----
-----------------
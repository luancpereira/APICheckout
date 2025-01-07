-----------------
---- INSERTS ----
-----------------

-- name: InsertTransaction :one
INSERT INTO "order" (
    description,
    transaction_date,
    transaction_value
) VALUES (
    @description::VARCHAR,
    @transaction_date::TIMESTAMP,
    @transaction_value::FLOAT
) RETURNING id;

-----------------
---- INSERTS ----
-----------------

-----------------
---- SELECTS ----
-----------------

-- name: SelectTransactions :many
SELECT 
    id,
    description,
    transaction_date::TIMESTAMP AS transaction_date,
    transaction_value
FROM
    "order";

-- name: SelectTransactionsTotal :one
SELECT 
    count(id) AS total
FROM
    "order";

-----------------
---- SELECTS ----
-----------------
// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: checkout.sql

package sqlc

import (
	"context"
	"time"
)

const deleteTransactionByID = `-- name: DeleteTransactionByID :exec


DELETE FROM "order"
WHERE id = $1::INTEGER
`

// ---------------
// -- INSERTS ----
// ---------------
// ---------------
// -- DELETES ----
// ---------------
func (q *Queries) DeleteTransactionByID(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteTransactionByIDStmt, deleteTransactionByID, id)
	return err
}

const insertTransaction = `-- name: InsertTransaction :one

INSERT INTO "order" (
    description,
    transaction_date,
    transaction_value
) VALUES (
    $1::VARCHAR,
    $2::TIMESTAMP,
    $3::FLOAT
) RETURNING id
`

type InsertTransactionParams struct {
	Description      string
	TransactionDate  time.Time
	TransactionValue float64
}

// ---------------
// -- INSERTS ----
// ---------------
func (q *Queries) InsertTransaction(ctx context.Context, arg InsertTransactionParams) (int64, error) {
	row := q.queryRow(ctx, q.insertTransactionStmt, insertTransaction, arg.Description, arg.TransactionDate, arg.TransactionValue)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const selectTransactionByID = `-- name: SelectTransactionByID :one
SELECT 
    id,
	description,
    transaction_date::TIMESTAMP AS transaction_date,
    transaction_value
FROM 
	"order"
WHERE
	id = $1::BIGINT
`

type SelectTransactionByIDRow struct {
	ID               int64
	Description      string
	TransactionDate  time.Time
	TransactionValue float64
}

func (q *Queries) SelectTransactionByID(ctx context.Context, id int64) (SelectTransactionByIDRow, error) {
	row := q.queryRow(ctx, q.selectTransactionByIDStmt, selectTransactionByID, id)
	var i SelectTransactionByIDRow
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.TransactionDate,
		&i.TransactionValue,
	)
	return i, err
}

const selectTransactions = `-- name: SelectTransactions :many


SELECT 
    id,
    description,
    transaction_date::TIMESTAMP AS transaction_date,
    transaction_value
FROM
    "order"
WHERE
	(CASE WHEN $3::VARCHAR <> '' THEN transaction_date::DATE >= $3::DATE ELSE TRUE END)
    AND (CASE WHEN $3::VARCHAR <> '' THEN transaction_date::DATE <= $3::DATE ELSE TRUE END)
LIMIT $1::BIGINT
OFFSET $2::BIGINT
`

type SelectTransactionsParams struct {
	Column1         int64
	Column2         int64
	TransactionDate string
}

type SelectTransactionsRow struct {
	ID               int64
	Description      string
	TransactionDate  time.Time
	TransactionValue float64
}

// ---------------
// -- UPDATES ----
// ---------------
// ---------------
// -- SELECTS ----
// ---------------
func (q *Queries) SelectTransactions(ctx context.Context, arg SelectTransactionsParams) ([]SelectTransactionsRow, error) {
	rows, err := q.query(ctx, q.selectTransactionsStmt, selectTransactions, arg.Column1, arg.Column2, arg.TransactionDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SelectTransactionsRow{}
	for rows.Next() {
		var i SelectTransactionsRow
		if err := rows.Scan(
			&i.ID,
			&i.Description,
			&i.TransactionDate,
			&i.TransactionValue,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectTransactionsTotal = `-- name: SelectTransactionsTotal :one
SELECT 
    count(id) AS total
FROM
    "order"
WHERE
	(CASE WHEN $1::VARCHAR <> '' THEN transaction_date::DATE >= $1::DATE ELSE TRUE END)
    AND (CASE WHEN $1::VARCHAR <> '' THEN transaction_date::DATE <= $1::DATE ELSE TRUE END)
`

func (q *Queries) SelectTransactionsTotal(ctx context.Context, transactionDate string) (int64, error) {
	row := q.queryRow(ctx, q.selectTransactionsTotalStmt, selectTransactionsTotal, transactionDate)
	var total int64
	err := row.Scan(&total)
	return total, err
}

const updateTransaction = `-- name: UpdateTransaction :exec


UPDATE "order"
SET
    description = $1::VARCHAR,
    transaction_date = $2::TIMESTAMP,
    transaction_value = $3::FLOAT
WHERE
    id = $4::INTEGER
`

type UpdateTransactionParams struct {
	Description      string
	TransactionDate  time.Time
	TransactionValue float64
	ID               int32
}

// ---------------
// -- DELETES ----
// ---------------
// ---------------
// -- UPDATES ----
// ---------------
func (q *Queries) UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) error {
	_, err := q.exec(ctx, q.updateTransactionStmt, updateTransaction,
		arg.Description,
		arg.TransactionDate,
		arg.TransactionValue,
		arg.ID,
	)
	return err
}

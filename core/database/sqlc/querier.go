// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package sqlc

import (
	"context"
)

type Querier interface {
	//---------------
	//-- INSERTS ----
	//---------------
	InsertTransaction(ctx context.Context, arg InsertTransactionParams) (int64, error)
	SelectTransactionByID(ctx context.Context, id int64) (SelectTransactionByIDRow, error)
	//---------------
	//-- INSERTS ----
	//---------------
	//---------------
	//-- SELECTS ----
	//---------------
	SelectTransactions(ctx context.Context, arg SelectTransactionsParams) ([]SelectTransactionsRow, error)
	SelectTransactionsTotal(ctx context.Context, arg SelectTransactionsTotalParams) (int64, error)
}

var _ Querier = (*Queries)(nil)
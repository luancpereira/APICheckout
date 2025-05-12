package database

import (
	"github.com/luancpereira/APICheckout/core/database/sqlc"
	coreError "github.com/luancpereira/APICheckout/core/errors"
)

type Utils struct{}

func (Utils) CoreErrorDatabase(err error) *coreError.CoreError {
	return coreError.New("error.database", err.Error())
}

func TransactionReturnOneObject[T any](fn func(querier sqlc.Querier) (fnObj T, fnErr error)) (obj T, err error) {
	tx, err := CONN.Begin()
	if err != nil {
		err = coreError.New("error.database.transaction.begin", err.Error())
		return
	}
	defer tx.Rollback()

	qtx := sqlc.New(CONN).WithTx(tx)
	defer qtx.Close()

	obj, err = fn(qtx)
	if err != nil {
		err = coreError.New("error.database.transaction.execute", err.Error())
		return
	}

	err = tx.Commit()
	if err != nil {
		err = coreError.New("error.database.transaction.commit", err.Error())
		return
	}

	return
}

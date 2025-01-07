package service

import (
	"context"
	"math"
	"time"

	"github.com/luancpereira/APICheckout/core/database"
	"github.com/luancpereira/APICheckout/core/database/sqlc"
	coreError "github.com/luancpereira/APICheckout/core/errors"
)

type Checkout struct{}

/*****
funcs for creations
******/

func (c Checkout) CreateTransaction(description string, transaction_date time.Time, transaction_value float64) (ID int64, err error) {

	err = c.validateDescription(description)
	if err != nil {
		return
	}

	err = c.validateTrasactionValue(transaction_value)
	if err != nil {
		return
	}

	if math.Round(transaction_value*100) != transaction_value*100 {
		transaction_value = math.Round(transaction_value*100) / 100
	}

	params := sqlc.InsertTransactionParams{
		Description:      description,
		TransactionDate:  transaction_date,
		TransactionValue: transaction_value,
	}

	ID, err = database.DB_QUERIER.InsertTransaction(context.Background(), params)
	if err != nil {
		err = database.Utils{}.CoreErrorDatabase(err)
		return
	}

	return
}

/*****
funcs for creations
******/

/*****
funcs for gets
******/

func (Checkout) GetList() (models []sqlc.SelectTransactionsRow, total int64, err error) {

	models, err = database.DB_QUERIER.SelectTransactions(context.Background())
	if err != nil {
		err = database.Utils{}.CoreErrorDatabase(err)
		return
	}

	total, err = database.DB_QUERIER.SelectTransactionsTotal(context.Background())
	if err != nil {
		err = database.Utils{}.CoreErrorDatabase(err)
		return
	}

	return
}

/*****
funcs for gets
******/

/*****
funcs for validations
******/

func (Checkout) validateDescription(description string) (err error) {
	if len(description) == 0 {
		err = coreError.New("error.description.empty")
		return
	}

	if len(description) > 50 {
		err = coreError.New("error.description.too.long")
		return
	}

	return
}

func (Checkout) validateTrasactionValue(value float64) (err error) {
	if value <= 0 {
		err = coreError.New("error.value.not.positive")
		return
	}

	return
}

/*****
funcs for validations
******/

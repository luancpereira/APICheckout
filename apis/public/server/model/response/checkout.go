package response

import (
	"time"
)

/*****
struct for posts
******/

type ReturnInsertTransaction struct {
	ID int64 `json:"id"`
}

/*****
struct for posts
******/

/*****
struct for gets
******/

type GetOrders struct {
	ID               int64     `json:"id"`
	Description      string    `json:"description"`
	TransactionDate  time.Time `json:"transaction_date"`
	TransactionValue float64   `json:"transaction_value"`
}

/*****
struct for gets
******/

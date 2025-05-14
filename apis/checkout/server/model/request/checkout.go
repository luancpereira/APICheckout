package request

import "time"

/*****
struct for posts
******/

type InsertTransaction struct {
	Description      string    `json:"description"`
	TransactionDate  time.Time `json:"transaction_date"`
	TransactionValue float64   `json:"transaction_value"`
}

type PutTransaction struct {
	ID               int64     `json:"id"`
	Description      string    `json:"description"`
	TransactionDate  time.Time `json:"transaction_date"`
	TransactionValue float64   `json:"transaction_value"`
}

/*****
struct for posts
******/

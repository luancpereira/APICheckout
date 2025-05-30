// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package sqlc

import (
	"database/sql"
	"time"
)

type Order struct {
	ID               int64
	Description      string
	TransactionDate  time.Time
	TransactionValue float64
}

type User struct {
	ID                              int64
	Email                           string
	Name                            sql.NullString
	Password                        sql.NullString
	Permission                      sql.NullString
	CreatedAt                       time.Time
	TokenConfirmation               sql.NullString
	TokenConfirmationExpirationDate sql.NullTime
}

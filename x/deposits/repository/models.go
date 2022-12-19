// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package repository

import (
	"database/sql"
)

type Deposit struct {
	ID     int32
	TxHash string
	Wallet string
	Amount string
	Height int64
}

type Height struct {
	Height int64
}

type Wallet struct {
	ID      int32
	Wallet  string
	Balance string
	Spent   string
}

type Withdrawal struct {
	ID            int32
	CorrelationID string
	WalletID      int32
	Amount        string
	Fee           string
	Expiry        int64
	TxHash        sql.NullString
}
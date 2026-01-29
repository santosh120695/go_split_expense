package model

import "gorm.io/gorm"

type UserTransaction struct {
	gorm.Model
	UserId        int64   `json:"user_id"`
	TransactionId int64   `json:"transaction_id"`
	Share         float64 `json:"share"`
	NetBalance    float64 `json:"net_balance"`
}

package models

import "gorm.io/gorm"

type UserTransaction struct {
	gorm.Model
	UserId        int64   `json:"user_id"`
	TransactionId int64   `json:"transaction_id"`
	Share         float64 `json:"share"`
}

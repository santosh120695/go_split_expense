package model

import "gorm.io/gorm"

type UserTransaction struct {
	gorm.Model
	UserId        float64 `json:"user_id"`
	User          User    `json:"user"`
	TransactionId int64   `json:"transaction_id"`
	Share         float64 `json:"share"`
	NetBalance    float64 `json:"net_balance"`
}

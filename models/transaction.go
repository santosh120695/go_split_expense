package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Title    string  `json:"title"`
	Amount   float64 `json:"mount"`
	PaidById int64   `json:"paid_by_id"`
	GroupId  int64   `json:"group_id"`
	Users    []User  `json:"users" gorm:"many2many:user_transactions"`
}

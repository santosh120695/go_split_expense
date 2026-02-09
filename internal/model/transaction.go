package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	PaidById  float64   `json:"paid_by_id"`
	PaidBy    User      `json:"paid_by"`
	GroupId   int64     `json:"group_id"`
	Users     []User    `json:"users" gorm:"many2many:user_transactions"`
	CreatedAt time.Time `json:"created_at"`
}

func (transaction *Transaction) BeforeCreate(db *gorm.DB) {
	var group Group
	db.First(&group, "id = ?", transaction.GroupId)
	group.TotalAmount += transaction.Amount
	db.Save(&group)
}

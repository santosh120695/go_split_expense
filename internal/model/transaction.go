package model

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	PaidById  float64   `json:"paid_by_id"`
	PaidBy    User      `json:"paid_by"`
	GroupId   float64   `json:"group_id"`
	Users     []User    `json:"users" gorm:"many2many:user_transactions"`
	CreatedAt time.Time `json:"created_at"`
}

func (transaction *Transaction) BeforeCreate(db *gorm.DB) (err error) {
	var group Group
	db.First(&group, "id = ?", transaction.GroupId)
	group.TotalAmount += transaction.Amount
	err = db.Save(&group).Error
	return err
}

func (transaction *Transaction) AfterCreate(db *gorm.DB) (err error) {
	db.Preload("Users").Find(&transaction.Users)
	activity := Activity{
		Description: transaction.PaidBy.UserName + " Paid " + strconv.FormatFloat(transaction.Amount, 'f', 2, 64) + " for " + transaction.Title,
		UserId:      transaction.PaidById,
		GroupId:     transaction.GroupId,
	}

	err = db.Create(&activity).Error
	return err
}

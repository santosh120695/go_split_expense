package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName          string `json:"user_name"`
	Email             string `gorm:"uniq;uniqueIndex;not null" json:"email"`
	EncryptedPassword string
	ContactNo         string        `json:"contact_no"`
	Groups            []Group       `json:"groups" gorm:"many2many:user_groups"`
	PaidTransactions  []Transaction `json:"paid_transactions" gorm:"foreignKey:PaidById"`
	Transactions      []Transaction `json:"transactions" gorm:"many2many:user_transactions"`
}

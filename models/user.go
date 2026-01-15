package models

import "time"

type User struct {
	ID                int64  `json:"id" gorm:"autoIncrement"`
	UserName          string `json:"user_name"`
	Email             string `gorm:"uniq;uniqueIndex;not null" json:"email"`
	EncryptedPassword string
	ContactNo         string        `json:"contact_no"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	Groups            []Group       `json:"groups" gorm:"many2many:user_groups"`
	PaidTransactions  []Transaction `json:"paid_transactions" gorm:"foreignKey:PaidById"`
	Transactions      []Transaction `json:"transactions" gorm:"many2many:user_transactions"`
}

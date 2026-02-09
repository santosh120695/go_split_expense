package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name         string        `json:"name"`
	CreatedById  float64       `json:"created_by_id"`
	Users        []User        `gorm:"many2many:user_groups"`
	Transactions []Transaction `json:"transactions"`
	Closed       bool          `gorm:"default:false" json:"closed"`
	Currency     string        `json:"currency"`
	Icon         string        `gorm:"default:Users" json:"icon"`
	TotalAmount  float64       `json:"total_amount" gorm:"default:0"`
	Description  string        `json:"description"`
}

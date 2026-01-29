package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name         string        `json:"name"`
	CreatedById  int64         `json:"created_by_id"`
	Users        []User        `gorm:"many2many:user_groups"`
	Transactions []Transaction `json:"transactions"`
}

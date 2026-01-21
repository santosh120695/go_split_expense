package models

import "time"

type Group struct {
	Id           int64         `json:"id" gorm:"autoIncrement"`
	Name         string        `json:"name"`
	CreatedById  int64         `json:"created_by_id"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	Users        []User        `gorm:"many2many:user_groups"`
	Transactions []Transaction `json:"transactions"`
}

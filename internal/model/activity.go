package model

import "gorm.io/gorm"

type Activity struct {
	gorm.Model  `json:"gorm_._model"`
	Description string  `json:"description" json:"description,omitempty"`
	UserId      float64 `json:"user_id,omitempty"`
	User        User    `gorm:"foreignkey:user_id" json:"user" json:"user"`
	GroupId     float64 `json:"group_id,omitempty"`
	Group       Group   `gorm:"foreignkey:group_id" json:"group,omitempty"`
}

package model

import (
	"gorm.io/gorm"
)

type UserGroup struct {
	gorm.Model
	UserId     float64
	GroupId    float64
	AmountOwe  float64
	AmountPaid float64
	User       User
}

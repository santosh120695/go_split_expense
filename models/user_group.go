package models

import (
	"gorm.io/gorm"
)

type UserGroup struct {
	gorm.Model
	UserId     int64
	GroupId    int64
	AmountOwe  float64
	AmountPaid float64
}

package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DashboardIndex(c *gin.Context, db *gorm.DB) {
	var activities []string
	userId, _ := c.Get("current_user")

	db.Raw("SELECT description FROM activities WHERE user_id = ? ORDER BY created_at DESC", userId).Scan(&activities)
	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"you_are_owed": gin.H{
				"amount":   300,
				"currency": "₹",
			},
			"you_owe": gin.H{
				"amount":   1000,
				"currency": "₹",
			},
			"activities": activities,
		},
	})

}

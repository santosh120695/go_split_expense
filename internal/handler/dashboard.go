package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DashboardIndex(c *gin.Context, db *gorm.DB) {
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
			"activities": []string{
				"You paid $50 to John for dinner",
				"Sarah paid you $25 for coffee",
				"New expense added: Groceries - $85.30",
				"Group 'House Bills' settlement completed",
				"You joined group \"Vacation Trip 2024\"",
			},
		},
	})

}

package handlers

import (
	"fmt"
	"net/http"
	"splitwise/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionIndexParams struct {
	GroupId int64 `form:"group_id"`
}

func TransactionIndex(c *gin.Context, db *gorm.DB) {
	var params TransactionIndexParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	var transactions []models.Transaction

	query := db
	fmt.Println(params.GroupId)
	if params.GroupId != 0 {
		query = query.Where("group_id = ?", params.GroupId).Find(&transactions)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactions,
	})
}

type CreateTransactionParams struct {
	GroupID  int64   `json:"group_id"`
	PaidById int64   `json:"paid_by_id"`
	Amount   float64 `json:"amount"`
	Title    string  `json:"title"`
	UserIds  []int64 `json:"user_ids"`
}

func TransactionCreate(c *gin.Context, db *gorm.DB) {
	var params CreateTransactionParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	var users []models.User
	db.Find(&users, params.UserIds)
	transaction := models.Transaction{
		Amount:   params.Amount,
		GroupId:  params.GroupID,
		PaidById: params.PaidById,
		Title:    params.Title,
		Users:    users,
	}

	fmt.Println(transaction)

	db.Save(&transaction)

	c.JSON(http.StatusOK, gin.H{
		"data": transaction,
	})
}

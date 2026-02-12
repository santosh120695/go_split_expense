package handler

import (
	"net/http"
	"splitwise/internal/model"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userSearchParams struct {
	SearchTerm string `form:"search_term"`
}

type UserSearchResponse struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	ID       float64
}

func UserSearch(c *gin.Context, db *gorm.DB) {
	searchParams := userSearchParams{}

	if err := c.ShouldBindQuery(&searchParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var users []UserSearchResponse
	term := strings.TrimSpace(searchParams.SearchTerm)

	if len(term) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "search term must be at least two characters"})
		return
	}

	limit := 10
	db.WithContext(c.Request.Context()).Raw("SELECT user_name, id, email FROM users WHERE user_name LIKE ? LIMIT ?", "%"+term+"%", limit).Scan(&users)
	c.JSON(http.StatusOK, gin.H{
		"data":  users,
		"count": len(users),
	})
}

func UserDetail(c *gin.Context, db *gorm.DB) {
	userId, _ := c.Get("current_user")
	var user model.User
	err := db.WithContext(c.Request.Context()).Where("id = ?", userId).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

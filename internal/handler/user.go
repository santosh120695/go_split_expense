package handler

import (
	"net/http"
	"splitwise/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userSearchParams struct {
	SearchTerm string `json:"search_term"`
}

type UserSearchResponse struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	ID       string `json:"id"`
}

func UserSearch(c *gin.Context, db *gorm.DB) {
	search_params := userSearchParams{}

	if err := c.ShouldBindJSON(&search_params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	var users []UserSearchResponse

	// db.Select("email", "user_name", "id").Where("user_name LIKE ?", "%"+search_params.SearchTerm+"%").Scan(&users)
	db.Model(&model.User{}).Select("email", "user_name", "id").Where("user_name LIKE ?", "%"+search_params.SearchTerm+"%").Scan(&users).Select("user_name", "email")

	c.JSON(http.StatusOK, gin.H{
		"data":  users,
		"count": len(users),
	})
}

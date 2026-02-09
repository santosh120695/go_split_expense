package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userSearchParams struct {
	SearchTerm string `form:"search_term"`
}

type UserSearchResponse struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	ID       string `json:"id"`
}

func UserSearch(c *gin.Context, db *gorm.DB) {
	searchParams := userSearchParams{}

	if err := c.ShouldBindQuery(&searchParams); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var users []UserSearchResponse
	fmt.Println(searchParams.SearchTerm)
	db.Raw("SELECT user_name, id, email FROM users WHERE user_name LIKE ?", "%"+searchParams.SearchTerm+"%").Scan(&users)
	fmt.Println(users)
	//db.Model(&model.User{}).Select("email", "user_name", "id").Where("user_name LIKE ?", "%"+searchParams.SearchTerm+"%").Scan(&users)

	c.JSON(http.StatusOK, gin.H{
		"data":  users,
		"count": len(users),
	})
}

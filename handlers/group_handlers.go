package handlers

import (
	"net/http"
	"splitwise/models"

	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GroupIndex(c *gin.Context, db *gorm.DB) {
	var groups models.Group

	db.Preload(clause.Associations).Find(&groups)

	c.JSON(http.StatusOK, gin.H{
		"data": groups,
	})
}

type CreateGroupParams struct {
	Name string `json:"name"`
}

func GroupCreate(c *gin.Context, db *gorm.DB) {
	var params CreateGroupParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	fmt.Println(params.Name)

	group := models.Group{Name: params.Name}

	db.Create(&group)

	c.JSON(http.StatusOK, gin.H{
		"data":    group,
		"success": true,
	})
}

type UpdateGroupParams struct {
	Name string `json:"name"`
	ID   int64  `uri:"id"`
}

func GroupUpdate(c *gin.Context, db *gorm.DB) {
	var params UpdateGroupParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}

	var group models.Group

	db.Where("id = ?", params.ID).First(&group)
	group.Name = params.Name
	db.Save(&group)

	c.JSON(http.StatusOK, gin.H{
		"data": group,
	})
}

func GroupDestroy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

type AdduserParams struct {
	UserIds []int64 `json:"user_ids"`
	ID      int64   `uri:"id"`
}

func AddUser(c *gin.Context, db *gorm.DB) {
	var params AdduserParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	var users []models.User

	var group models.Group

	db.Where("id = ?", params.ID).First(&group)

	db.Find(&users, params.UserIds)

	group.Users = users
	db.Save(&group)

	c.JSON(http.StatusOK, gin.H{
		"groups":     group,
		"user_count": len(group.Users),
	})
}

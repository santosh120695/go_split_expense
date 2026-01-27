package handler

import (
	"fmt"
	"net/http"
	"splitwise/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GroupIndex(c *gin.Context, db *gorm.DB) {
	var groups model.Group

	db.Preload(clause.Associations).Find(&groups)

	c.JSON(http.StatusOK, gin.H{
		"data": groups,
	})
}

type GroupDetail struct {
	UserName string  `json:"user_name"`
	Email    string  `json:"email"`
	UserId   int64   `json:"user_id"`
	Pay      float64 `json:"pay"`
	Receive  float64 `json:"receive"`
}

type GroupShowParam struct {
	ID int64 `uri:"id"`
}

func GroupShow(c *gin.Context, db *gorm.DB) {

	var group model.Group
	var user_groups []model.UserGroup
	var group_details []GroupDetail
	var params GroupShowParam

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	db.Preload(clause.Associations).Where("id = ? ", params.ID).Find(&group)
	db.Preload(clause.Associations).Where("group_id = ?", group.ID).Find(&user_groups)

	for _, user_group := range user_groups {
		fmt.Println(user_group.AmountOwe)
		fmt.Println(user_group.AmountPaid)
		group_details = append(group_details, GroupDetail{
			UserName: user_group.User.UserName,
			Email:    user_group.User.Email,
			UserId:   user_group.UserId,
			Pay:      user_group.AmountOwe - user_group.AmountPaid,
			Receive:  user_group.AmountPaid - user_group.AmountOwe,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"name":  group.Name,
		"users": group_details,
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

	group := model.Group{Name: params.Name}

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

	var group model.Group

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

	var users []model.User

	var group model.Group

	db.Where("id = ?", params.ID).First(&group)

	db.Find(&users, params.UserIds)

	group.Users = users
	db.Save(&group)

	c.JSON(http.StatusOK, gin.H{
		"groups":     group,
		"user_count": len(group.Users),
	})
}

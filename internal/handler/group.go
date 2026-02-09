package handler

import (
	"fmt"
	"net/http"
	"splitwise/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GroupIndexResponse struct {
	Name         string       `json:"name"`
	Icon         string       `json:"icon"`
	Description  string       `json:"description"`
	Members      []model.User `json:"members"`
	TotalExpense float64      `json:"total_expense"`
	Currency     string       `json:"currency"`
	ID           uint         `json:"id"`
}

func GroupIndex(c *gin.Context, db *gorm.DB) {
	userId, _ := c.Get("current_user")
	var user model.User
	var response []GroupIndexResponse
	db.Preload(clause.Associations).First(&user, userId.(float64))
	db.Preload(clause.Associations).Find(&user.Groups)
	for _, group := range user.Groups {
		totalExpense := 0.0
		for _, transaction := range group.Transactions {
			totalExpense += transaction.Amount
		}
		response = append(response, GroupIndexResponse{
			Name:         group.Name,
			Icon:         group.Icon,
			Description:  "",
			Members:      group.Users,
			TotalExpense: totalExpense,
			Currency:     group.Currency,
			ID:           group.ID,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    response,
		"success": true,
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
	var userGroups []model.UserGroup
	var groupDetails []GroupDetail
	var params GroupShowParam

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	fmt.Println(params.ID)

	db.Preload(clause.Associations).Where("id = ? ", params.ID).Find(&group)
	db.Preload(clause.Associations).Where("group_id = ?", group.ID).Find(&userGroups)
	db.Preload(clause.Associations).Find(&group.Transactions)

	for _, userGroup := range userGroups {
		groupDetails = append(groupDetails, GroupDetail{
			UserName: userGroup.User.UserName,
			Email:    userGroup.User.Email,
			UserId:   int64(userGroup.UserId),
			Pay:      userGroup.AmountOwe - userGroup.AmountPaid,
			Receive:  userGroup.AmountPaid - userGroup.AmountOwe,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"name":        group.Name,
			"description": group.Description,
			"members":     groupDetails,
			"expenses":    group.Transactions,
		},
	})
}

type CreateGroupParams struct {
	Name        string `json:"name"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
}

func GroupCreate(c *gin.Context, db *gorm.DB) {
	var params CreateGroupParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	userId, _ := c.Get("current_user")
	group := model.Group{Name: params.Name, CreatedById: userId.(float64), Currency: params.Currency}

	db.Create(&group)
	userGroup := model.UserGroup{GroupId: float64(group.ID), UserId: userId.(float64)}
	db.Create(&userGroup)

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

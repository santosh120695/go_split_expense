package handler

import (
	"net/http"
	"slices"
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
	ID           uint
}

func GroupIndex(c *gin.Context, db *gorm.DB) {
	userId, _ := c.Get("current_user")
	var user model.User
	var response []GroupIndexResponse

	db.WithContext(c.Request.Context()).Preload(clause.Associations).First(&user, userId.(float64))
	db.WithContext(c.Request.Context()).Preload(clause.Associations).Find(&user.Groups)

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
	var params GroupShowParam

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.WithContext(c.Request.Context()).Preload(clause.Associations).Where("id = ? ", params.ID).Find(&group)
	db.WithContext(c.Request.Context()).Preload(clause.Associations).Find(&group.Transactions, "group_id = ? ", group.ID)

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"name":        group.Name,
			"description": group.Description,
			"members":     group.Users,
			"expenses":    group.Transactions,
			"activities":  group.Activities,
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userId, _ := c.Get("current_user")
	group := model.Group{Name: params.Name, CreatedById: userId.(float64), Currency: params.Currency}

	db.WithContext(c.Request.Context()).Create(&group)
	userGroup := model.UserGroup{GroupId: float64(group.ID), UserId: userId.(float64)}
	go db.Create(&userGroup)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    group,
	})
}

type UpdateGroupParams struct {
	Name string `json:"name"`
	ID   int64  `uri:"id"`
}

func GroupUpdate(c *gin.Context, db *gorm.DB) {
	var params UpdateGroupParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var group model.Group

	db.WithContext(c.Request.Context()).Where("id = ?", params.ID).First(&group)
	group.Name = params.Name
	db.WithContext(c.Request.Context()).Save(&group)

	c.JSON(http.StatusOK, gin.H{
		"data": group,
	})
}

type AdduserParams struct {
	UserIds []float64 `json:"user_ids"`
	ID      float64   `uri:"id"`
}

func AddUser(c *gin.Context, db *gorm.DB) {
	var params AdduserParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	var userGroups []model.UserGroup

	var usersAlreadyAdded []float64

	db.Model(model.UserGroup{}).Where("group_id = ?", params.ID).Select("user_id").Find(&usersAlreadyAdded)

	for _, userId := range params.UserIds {
		if !slices.Contains(usersAlreadyAdded, userId) {
			userGroups = append(userGroups, model.UserGroup{
				UserId:  userId,
				GroupId: params.ID,
			})
		}
	}

	db.WithContext(c.Request.Context()).Save(&userGroups)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

type GroupRepayParams struct {
	GroupId []float64 `uri:"id"`
}

func GroupRepays(c *gin.Context, db *gorm.DB) {
	var params GroupRepayParams
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	var group model.Group

	db.WithContext(c.Request.Context()).First(&group, "id = ?", params.GroupId)

	repays := group.CalculateRepayments(db)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   repays,
	})
}

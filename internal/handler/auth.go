package handler

import (
	"errors"
	"net/http"
	"splitwise/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SignUpParams struct {
	UserName        string `json:"user_name"`
	ContactNo       string `json:"contact_no"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Email           string `json:"email"`
}

func SignUp(c *gin.Context, db *gorm.DB) {
	signupParam := SignUpParams{}
	if err := c.ShouldBindJSON(&signupParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if signupParam.ConfirmPassword != signupParam.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "confirm_password and password should be same.",
		})
		return
	}

	user := model.User{UserName: signupParam.UserName, Email: signupParam.Email, Password: signupParam.Password, ContactNo: signupParam.ContactNo}

	result := db.WithContext(c.Request.Context()).Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"email":     signupParam.Email,
		"user_name": signupParam.UserName,
	})
}

type SignParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignIn(c *gin.Context, db *gorm.DB) {
	signingParams := SignParams{}
	if err := c.ShouldBindJSON(&signingParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	var user model.User
	//db.Where("email = ?", signingParams.Email).First(&user)
	result := db.WithContext(c.Request.Context()).First(&user, "email = ?", signingParams.Email)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"status": false, "error": "Invalid email or password"})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "error": result.Error.Error()})
	}
	if user.Email == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "user not found",
		})
		return
	}

	if user.VerifyPassword(signingParams.Password) {
		token, err := user.CreateToken()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"token":   token,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"message": "Invalid email or password!",
	})
}

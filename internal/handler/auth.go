package handler

import (
	"net/http"
	"os"
	"splitwise/internal/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SignUpParams struct {
	UserName        string `json:"user_name"`
	ContactNo       string `json:"contact_no"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Email           string `json:"email"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createToken(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	var secretKey = []byte(os.Getenv("SECRET_KEY"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func SignUp(c *gin.Context, db *gorm.DB) {
	signupParam := SignUpParams{}
	if err := c.ShouldBindJSON(&signupParam); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if signupParam.ConfirmPassword != signupParam.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "confirm_password and password should be same.",
		})
		return
	}

	encryptedPassword, err := hashPassword(signupParam.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := model.User{UserName: signupParam.UserName, Email: signupParam.Email, EncryptedPassword: encryptedPassword, ContactNo: signupParam.ContactNo}

	result := db.Create(&user)
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
	db.Where("email = ?", signingParams.Email).First(&user)

	if user.Email == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "user not found",
		})
		return
	}
	if verifyPassword(signingParams.Password, user.EncryptedPassword) {
		token, err := createToken(user)

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

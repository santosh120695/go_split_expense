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
	UserName  string `json:"user_name"`
	ContactNo string `json:"contact_no"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
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
	signup_param := SignUpParams{}

	if err := c.ShouldBindJSON(&signup_param); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	encrypted_password, err := hashPassword(signup_param.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	user := model.User{UserName: signup_param.UserName, Email: signup_param.Email, EncryptedPassword: encrypted_password}

	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	}

	c.JSON(http.StatusOK, gin.H{
		"email":     signup_param.Email,
		"user_name": signup_param.UserName,
	})
}

type SignParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignIn(c *gin.Context, db *gorm.DB) {
	signin_params := SignParams{}
	if err := c.ShouldBindJSON(&signin_params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}

	var user model.User
	db.Where("email = ?", signin_params.Email).First(&user)

	if user.Email == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "user not found",
		})
	}

	if verifyPassword(signin_params.Password, user.EncryptedPassword) {
		token, err := createToken(user.Email)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"error":   err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"token":   token,
		})
	}
}

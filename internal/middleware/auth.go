package middleware

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization token",
			})
			return
		}
		token, err := verifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "invalid token",
			})
		}
		c.Set("current_user", token.Claims.(jwt.MapClaims)["id"].(float64))
		c.Next()
	}
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	var secretKey = []byte(os.Getenv("SECRET_KEY"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})
	if !token.Valid {
		return token, errors.New("invalid token")
	}
	return token, err
}

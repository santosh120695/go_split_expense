package configs

import (
	"errors"
	"splitwise/handlers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Route struct {
	RequestType  string
	Handler      func(*gin.Context, *gorm.DB)
	RequestGroup string
	Path         string
	Authenticate bool
}

var Routes = []Route{
	{RequestType: "POST", Handler: handlers.SignUp, RequestGroup: "users", Path: "/sign_up", Authenticate: false},
	{RequestType: "POST", Handler: handlers.SignIn, RequestGroup: "users", Path: "/sign_in", Authenticate: false},
	{RequestType: "GET", Handler: handlers.UserSearch, RequestGroup: "users", Path: "/search", Authenticate: true},
	{RequestType: "GET", Handler: handlers.GroupIndex, RequestGroup: "groups", Path: "", Authenticate: true},
	{RequestType: "GET", Handler: handlers.GroupShow, RequestGroup: "groups", Path: "/:id", Authenticate: true},
	{RequestType: "POST", Handler: handlers.GroupCreate, RequestGroup: "groups", Path: "", Authenticate: true},
	{RequestType: "PUT", Handler: handlers.GroupUpdate, RequestGroup: "groups", Path: "/:id", Authenticate: true},
	{RequestType: "POST", Handler: handlers.AddUser, RequestGroup: "groups", Path: "/:id/add_users", Authenticate: true},
	{RequestType: "GET", Handler: handlers.TransactionIndex, RequestGroup: "transactions", Path: "", Authenticate: true},
	{RequestType: "POST", Handler: handlers.TransactionCreate, RequestGroup: "transactions", Path: "", Authenticate: true},
	{RequestType: "GET", Handler: handlers.CalculateSplit, RequestGroup: "transactions", Path: "/:group_id", Authenticate: true},
}

func InitRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	for _, route := range Routes {
		group := v1.Group(route.RequestGroup)
		if route.Authenticate {
			group.Use(AuthMiddleWare())
		}
		group.Handle(route.RequestType, route.Path, func(c *gin.Context) {
			route.Handler(c, db)
		})
	}
	return r
}

func verifyToken(tokenString string) error {
	var secretKey = []byte("some-key")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})
	if !token.Valid {
		return errors.New("invalid token")
	}
	return err
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "missing authorization header",
			})
		}
		if verifyToken(tokenString) != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"error": "invalid token",
			})
		}
		c.Next()
	}
}

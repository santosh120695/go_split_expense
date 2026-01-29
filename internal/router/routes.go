package router

import (
	"splitwise/internal/handler"
	"splitwise/internal/middleware"

	"github.com/gin-gonic/gin"
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
	{RequestType: "POST", Handler: handler.SignUp, RequestGroup: "users", Path: "/sign_up", Authenticate: false},
	{RequestType: "POST", Handler: handler.SignIn, RequestGroup: "users", Path: "/sign_in", Authenticate: false},
	{RequestType: "GET", Handler: handler.UserSearch, RequestGroup: "users", Path: "/search", Authenticate: true},
	{RequestType: "GET", Handler: handler.GroupIndex, RequestGroup: "groups", Path: "", Authenticate: true},
	{RequestType: "GET", Handler: handler.GroupShow, RequestGroup: "groups", Path: "/:id", Authenticate: true},
	{RequestType: "POST", Handler: handler.GroupCreate, RequestGroup: "groups", Path: "", Authenticate: true},
	{RequestType: "PUT", Handler: handler.GroupUpdate, RequestGroup: "groups", Path: "/:id", Authenticate: true},
	{RequestType: "POST", Handler: handler.AddUser, RequestGroup: "groups", Path: "/:id/add_users", Authenticate: true},
	{RequestType: "GET", Handler: handler.TransactionIndex, RequestGroup: "transactions", Path: "", Authenticate: true},
	{RequestType: "POST", Handler: handler.TransactionCreate, RequestGroup: "transactions", Path: "", Authenticate: true},
	{RequestType: "GET", Handler: handler.CalculateSplit, RequestGroup: "transactions", Path: "/:group_id", Authenticate: true},
}

func InitRoutes(db *gorm.DB, r *gin.Engine) {
	v1 := r.Group("/v1")
	for _, route := range Routes {
		group := v1.Group(route.RequestGroup)
		if route.Authenticate {
			group.Use(middleware.AuthMiddleWare())
		}
		group.Handle(route.RequestType, route.Path, func(c *gin.Context) {
			route.Handler(c, db)
		})
	}
}

package main

import (
	"fmt"
	"log"
	"splitwise/internal/config"
	"splitwise/internal/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("welcome to splitwise app")

	db := config.ConnectDB()
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: false,
	}))
	router.InitRoutes(db, engine)
	if err := engine.Run(":3000"); err != nil {
		log.Fatalf("failed to start server. %v", err)
	}
}

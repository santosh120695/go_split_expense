package main

import (
	"log"
	"os"
	"splitwise/internal/config"
	"splitwise/internal/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err.Error())
			return
		}
	}

	db := config.ConnectDB()
	gin.SetMode(gin.ReleaseMode)
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

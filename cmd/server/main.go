package main

import (
	"fmt"
	"log"
	"splitwise/internal/config"
	"splitwise/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("welcome to splitwise app")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error in loading env file")
	}
	db := config.ConnectDB()
	engine := router.InitRoutes(db)
	if err := engine.Run(":3000"); err != nil {
		log.Fatalf("failed to start server. %v", err)
	}
}

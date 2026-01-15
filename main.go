package main

import (
	"fmt"
	"log"
	"splitwise/configs"
)

func main() {
	fmt.Println("welcome to splitwise app")

	db := configs.ConnectDB()
	engine := configs.InitRoutes(db)
	if err := engine.Run(":3000"); err != nil {
		log.Fatalf("failed to start server. %v", err)
	}
}

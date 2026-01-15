package configs

import (
	"fmt"
	"splitwise/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./splitwise.db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("error in connecting with db, %v\n", err.Error())
	} else {
		db.AutoMigrate(&models.User{}, &models.Group{})
	}
	return db
}

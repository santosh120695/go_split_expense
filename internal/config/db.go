package config

import (
	"fmt"
	"os"
	"splitwise/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_NAME")), &gorm.Config{})
	if err != nil {
		fmt.Printf("error in connecting with db, %v\n", err.Error())
	} else {
		db.AutoMigrate(&model.User{},
			&model.Group{},
			&model.UserGroup{},
			&model.Transaction{},
			&model.UserTransaction{},
		)
	}
	return db
}

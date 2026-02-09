package config

import (
	"fmt"
	"os"
	"splitwise/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := openDB()
	if err != nil {
		fmt.Printf("error in connecting with db, %v\n", err.Error())
	} else {
		err := db.AutoMigrate(&model.User{},
			&model.Group{},
			&model.UserGroup{},
			&model.Transaction{},
			&model.UserTransaction{},
		)
		if err != nil {
			return nil
		}
	}
	return db
}

func openDB() (*gorm.DB, error) {
	env := os.Getenv("ENV")

	dbPath := os.Getenv("DB_NAME")
	if env == "production" {
		return gorm.Open(postgres.Open(dbPath), &gorm.Config{})
	}

	return gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
}

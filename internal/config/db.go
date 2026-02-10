package config

import (
	"context"
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
		err := db.WithContext(context.Background()).AutoMigrate(&model.User{},
			&model.Group{},
			&model.UserGroup{},
			&model.Transaction{},
			&model.UserTransaction{},
			&model.Activity{},
		)
		if err != nil {
			return nil
		}
		seedDB(db)
	}
	return db
}

func openDB() (*gorm.DB, error) {
	env := os.Getenv("ENV")

	dbPath := os.Getenv("DATABASE_URL")
	if env == "production" {
		return gorm.Open(postgres.Open(dbPath), &gorm.Config{})
	}

	return gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
}

func seedDB(db *gorm.DB) {
	var users []model.User
	users = append(users, model.User{
		UserName: "Santosh ghode",
		Password: "1Linkwok@",
		Email:    "ghodesantosh0@gmail.com",
	})

	users = append(users, model.User{
		UserName: "Vishal Patel	",
		Password: "1Linkwok@",
		Email:    "vishal.patel@gmail.com",
	})

	users = append(users, model.User{
		UserName: "Kiran manwar",
		Password: "1Linkwok@",
		Email:    "kiron.manwar@gmail.com",
	})

	err := db.Create(&users)
	if err != nil {
		fmt.Println(err.Error)
	}
}

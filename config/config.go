package config

import (
	"log"
	"main/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("uninacorsistudiobot.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	err = DB.AutoMigrate(&models.Department{}, &models.Course{}, &models.User{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	log.Println("Database initialized")
}

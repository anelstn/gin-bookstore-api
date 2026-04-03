package config

import (
	"fmt"
	"log"
	"os"

	"git-practice-gin/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	_ = godotenv.Load()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Almaty",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}
	err = DB.AutoMigrate(&models.Author{}, &models.Category{}, &models.Book{})
	if err != nil {
		log.Fatal("Failed to AutoMigrate:", err)
	}
	log.Println("Connected to PostgreSQL and migrated successfully!")
}

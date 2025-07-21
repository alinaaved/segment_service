package db

import (
	"fmt"
	"log"
	"os"

	"segment_service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=segment_service port=5432 sslmode=disable"
	}
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	database.AutoMigrate(&models.User{}, &models.Segment{})
	DB = database

	fmt.Println("Database connected and migrated")
}

func GetDB() *gorm.DB {
	return DB
}

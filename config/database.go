package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Databases initializes the database connection
func Databases() {
	App() // Load environment variables

	dsn := GetEnv("DB_USERNAME") + ":" + GetEnv("DB_PASSWORD") + "@tcp(" + GetEnv("DB_HOST") + ":" + GetEnv("DB_PORT") + ")/" + GetEnv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

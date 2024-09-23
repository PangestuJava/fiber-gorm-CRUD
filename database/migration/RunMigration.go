package migration

import (
	"fiber-gorm-CRUD/app/models"
	"fiber-gorm-CRUD/config"
	"log"
)

func RunMigration() {
	err := config.DB.AutoMigrate(&models.Categories{})

	if err != nil {
		log.Fatalf("Failed to migrate category: %v", err)
	}
}

package main

import (
	"fiber-gorm-CRUD/config"
	"fiber-gorm-CRUD/database/migration"
	"fiber-gorm-CRUD/routes"

	"github.com/gofiber/fiber/v2"

)

func main() {
	serve := fiber.New()

	//connect to database
	config.Databases()

	//run migration
	migration.RunMigration()

	// Call the function to set up guest routes (no middleware)
	routes.GuestRoutes(serve)

	// Call the function to set up protected API routes (with middleware)
	routes.ApiRoutes(serve)

	serve.Listen(":8000")
}

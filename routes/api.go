package routes

import (
	"fiber-gorm-CRUD/app/http/controllers/admins"

	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(serve *fiber.App) {
	// Group routes under /api prefix
	api := serve.Group("/api/admin/")

	// Define your routes here
	api.Get("categories", admins.CategoryIndex)
	api.Post("category", admins.CategoryCreate)
	api.Put("category/:id", admins.CategoryUpdate)
	api.Delete("category/:id", admins.CategoryDelete)

}

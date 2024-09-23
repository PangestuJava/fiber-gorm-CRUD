package routes

import (
	"runtime"

	"github.com/gofiber/fiber/v2"
)

func GuestRoutes(serve *fiber.App) {
	serve.Get("/", func(c *fiber.Ctx) error {
		goVersion := runtime.Version()
		fiberVersion := fiber.Version
		gormVersion := "v1.25.11"

		return c.JSON(fiber.Map{
			"go_version":    goVersion,
			"fiber_version": fiberVersion,
			"gorm_version":  gormVersion,
		})
	})

	// Group routes under /api prefix
	api := serve.Group("/api/public/")

	// Define guest routes here
	api.Get("get", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "This is a public route",
			"status":  "success",
		})
	})
}

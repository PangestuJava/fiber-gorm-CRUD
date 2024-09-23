package admins

import (
	"fiber-gorm-CRUD/app/http/requests/admins/category"
	"fiber-gorm-CRUD/app/models"
	"fiber-gorm-CRUD/app/traits"
	"fiber-gorm-CRUD/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

)

func CategoryIndex(ctx *fiber.Ctx) error {
	var categories []models.Categories

	paginatedData, err := traits.Paginate(ctx, config.DB, &categories, 2) // Pass pointer to slice

	if err != nil {
		return traits.JSONResponse(ctx, fiber.StatusInternalServerError, false, "Failed to retrieve categories", nil)
	}

	return traits.JSONResponse(ctx, fiber.StatusOK, true, "Categories retrieved successfully", paginatedData)
}

func CategoryCreate(ctx *fiber.Ctx) error {
	categoryRequest := new(category.CreateCategoryRequest)

	if err := ctx.BodyParser(categoryRequest); err != nil {
		return traits.JSONResponse(ctx, fiber.StatusBadRequest, false, "Failed to parse request body", nil)
	}

	validate := validator.New()
	if err := validate.Struct(categoryRequest); err != nil {
		return traits.ValidationError(ctx, err)
	}

	category := models.Categories{
		Name: categoryRequest.Name,
	}

	if err := config.DB.Create(&category).Error; err != nil {
		return traits.JSONResponse(ctx, fiber.StatusInternalServerError, false, "Failed to create category", nil)
	}

	return traits.JSONResponse(ctx, fiber.StatusCreated, true, "Category created successfully", category)
}

func CategoryUpdate(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var categoryRequest category.CreateCategoryRequest

	if err := ctx.BodyParser(&categoryRequest); err != nil {
		return traits.ValidationError(ctx, err)
	}

	var category models.Categories
	if err := config.DB.First(&category, id).Error; err != nil {
		return traits.NotFoundError(ctx, "Category not found")
	}

	category.Name = categoryRequest.Name

	if err := config.DB.Save(&category).Error; err != nil {
		return traits.DatabaseError(ctx, "Failed to update category")
	}

	return traits.JSONResponse(ctx, fiber.StatusOK, true, "Category updated successfully", category)
}

func CategoryDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var category models.Categories
	if err := config.DB.First(&category, id).Error; err != nil {
		return traits.NotFoundError(ctx, "Category not found")
	}

	if err := config.DB.Delete(&category).Error; err != nil {
		return traits.DatabaseError(ctx, "Failed to delete category")
	}

	return traits.JSONResponse(ctx, fiber.StatusOK, true, "Category deleted successfully", nil)
}

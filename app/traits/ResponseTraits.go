package traits

import (
	"log"
	"reflect"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Define response structures
type Response struct {
	Status  string      `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Pagination struct {
	CurrentPage int    `json:"current_page"`
	From        int    `json:"from"`
	LastPage    int    `json:"last_page"`
	Path        string `json:"path"`
	PerPage     int    `json:"per_page"`
	To          int    `json:"to"`
	Total       int    `json:"total"`
}

type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
}

type PaginatedData struct {
	Items      interface{} `json:"items"`
	Pagination Pagination  `json:"pagination"`
	Links      Links       `json:"links"`
}

func DatabaseError(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":  fiber.StatusInternalServerError,
		"success": false,
		"message": message,
	})
}

func ValidationError(ctx *fiber.Ctx, err error) error {
	validationErrors := err.(validator.ValidationErrors)
	errorsMap := make(map[string][]string)

	for _, err := range validationErrors {
		fieldName := err.Field()
		errorMessage := "The " + fieldName + " field is " + err.Tag() + "."

		if _, exists := errorsMap[fieldName]; !exists {
			errorsMap[fieldName] = []string{}
		}

		errorsMap[fieldName] = append(errorsMap[fieldName], errorMessage)
	}

	// Create a structured response to ensure key order
	response := struct {
		Message string              `json:"message"`
		Errors  map[string][]string `json:"errors"`
	}{
		Message: "Validation failed.",
		Errors:  errorsMap,
	}

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response)
}

func NotFoundError(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"status":  fiber.StatusNotFound,
		"success": false,
		"message": message,
	})
}

// Paginate function to handle pagination
func Paginate(ctx *fiber.Ctx, db *gorm.DB, model interface{}, perPage int) (PaginatedData, error) {
	var total int64
	page, _ := strconv.Atoi(ctx.Query("page", "1"))

	// Count total records
	db.Model(model).Count(&total)

	// Determine offset
	offset := (page - 1) * perPage

	// Fetch data with pagination
	err := db.Limit(perPage).Offset(offset).Find(model).Error
	if err != nil {
		return PaginatedData{}, err
	}

	// Calculate total pages
	lastPage := int((total + int64(perPage) - 1) / int64(perPage))

	// Use reflect to determine the length of the slice
	slice := reflect.ValueOf(model).Elem()
	to := offset + slice.Len()

	// Create pagination struct
	pagination := Pagination{
		CurrentPage: page,
		From:        offset + 1,
		LastPage:    lastPage,
		Path:        ctx.BaseURL() + ctx.Path(),
		PerPage:     perPage,
		To:          to,
		Total:       int(total),
	}

	// Create links
	links := Links{
		First: ctx.BaseURL() + ctx.Path() + "?page=1",
		Last:  ctx.BaseURL() + ctx.Path() + "?page=" + strconv.Itoa(lastPage),
	}

	if page > 1 {
		links.Prev = ctx.BaseURL() + ctx.Path() + "?page=" + strconv.Itoa(page-1)
	}
	if page < lastPage {
		links.Next = ctx.BaseURL() + ctx.Path() + "?page=" + strconv.Itoa(page+1)
	}

	return PaginatedData{
		Items:      model,
		Pagination: pagination,
		Links:      links,
	}, nil
}

// JSONResponse function to create a standardized JSON response
func JSONResponse(ctx *fiber.Ctx, statusCode int, success bool, message string, data interface{}) error {
	if !success {
		log.Println(message)
	}

	return ctx.Status(statusCode).JSON(Response{
		Status:  strconv.Itoa(statusCode),
		Success: success,
		Message: message,
		Data:    data,
	})
}

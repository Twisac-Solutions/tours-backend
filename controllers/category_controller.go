package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/gofiber/fiber/v2"
)

// GetAllCategories godoc
// @Summary      Get all categories
// @Description  Retrieves a list of all categories
// @Tags         categories
// @Produce      json
// @Success      200  {array}   models.Category
// @Failure      500  {object}  models.ErrorResponse
// @Router       /api/categories [get]
func GetAllCategories(c *fiber.Ctx) error {
	categories, err := services.GetAllCategories()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve categories"})
	}
	return c.JSON(categories)
}

// GetCategoryByID godoc
// @Summary      Get category by ID
// @Description  Retrieves a category by its ID
// @Tags         categories
// @Produce      json
// @Param        id   path      string  true  "Category ID"
// @Success      200  {object}  models.Category
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/categories/{id} [get]
func GetCategoryByID(c *fiber.Ctx) error {
	id := c.Params("id")
	category, err := services.GetCategoryByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Category not found"})
	}
	return c.JSON(category)
}

// CreateCategory godoc
// @Summary      Create a new category
// @Description  Creates a new category
// @Tags         admin_categories
// @Accept       json
// @Produce      json
// @Param        category  body      models.Category  true  "Category object"
// @Success      200   {object}  models.Category
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /admin/categories [post]
func CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Validate required fields
	if category.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Name is required"})
	}
	if category.Icon == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Icon is required"})
	}

	if err := services.CreateCategory(&category); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create category"})
	}
	return c.JSON(category)
}

// UpdateCategory godoc
// @Summary      Update a category
// @Description  Updates an existing category by ID
// @Tags         admin_categories
// @Accept       json
// @Produce      json
// @Param        id    path      string      true  "Category ID"
// @Param        category body      models.Category true  "Category object"
// @Param        coverImage formData file false "Cover image file"
// @Success      200   {object}  models.Category
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /admin/categories/{id} [put]
func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var updated models.Category
	if err := c.BodyParser(&updated); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := services.UpdateCategory(id, &updated); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update category"})
	}
	return c.JSON(updated)
}

// DeleteCategory godoc
// @Summary      Delete a category
// @Description  Deletes a category by ID
// @Tags         admin_categories
// @Produce      json
// @Param        id   path      string  true  "Category ID"
// @Success      200  {object}  models.MessageResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/categories/{id} [delete]
func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := services.DeleteCategory(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete category"})
	}
	return c.JSON(fiber.Map{"message": "Category deleted"})
}

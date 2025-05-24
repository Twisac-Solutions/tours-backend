package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAllCategories(c *fiber.Ctx) error {
	categories, err := services.GetAllCategories()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve categories"})
	}
	return c.JSON(categories)
}

func GetCategoryByID(c *fiber.Ctx) error {
	id := c.Params("id")
	category, err := services.GetCategoryByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Category not found"})
	}
	return c.JSON(category)
}

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	form, _ := c.MultipartForm()
	if form != nil {
		category.CoverImage.URL, _ = utils.SaveFile(form.File["coverImage"])
	}

	if err := services.CreateCategory(&category); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create category"})
	}
	return c.JSON(category)
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var updated models.Category
	if err := c.BodyParser(&updated); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	form, _ := c.MultipartForm()
	if form != nil {
		updated.CoverImage.URL, _ = utils.SaveFile(form.File["coverImage"])
	}

	if err := services.UpdateCategory(id, &updated); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update category"})
	}
	return c.JSON(updated)
}

func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := services.DeleteCategory(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete category"})
	}
	return c.JSON(fiber.Map{"message": "Category deleted"})
}

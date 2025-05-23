package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your-username/tout-api/services"
)

func GetUserProfile(c *fiber.Ctx) error {
	return services.GetUserProfile(c)
}

package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/gofiber/fiber/v2"
)

func GetUserProfile(c *fiber.Ctx) error {
	return services.GetUserProfile(c)
}

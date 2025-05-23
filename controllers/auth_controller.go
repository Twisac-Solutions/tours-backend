package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	return services.Register(c)
}

func Login(c *fiber.Ctx) error {
	return services.Login(c)
}

func GoogleSSO(c *fiber.Ctx) error {
	return services.GoogleSSO(c)
}

func Logout(c *fiber.Ctx) error {
	return services.Logout(c)
}

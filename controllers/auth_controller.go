package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your-username/tout-api/services"
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

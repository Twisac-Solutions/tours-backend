package middlewares

import (
	"log"

	"github.com/Twisac-Solutions/tours-backend/utils" // replace with your actual project import path
	"github.com/gofiber/fiber/v2"
)

func AdminOnly(c *fiber.Ctx) error {
	// Skip auth check for login route
	log.Println("Checking admin auth: ", c.Path())
	if c.Path() == "/admin/login" {
		return c.Next()
	}

	// Verify JWT and get role
	userID, role, err := utils.VerifyJWTRole(c)
	if err != nil {
		log.Printf("JWT verification failed: %v", err)
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Set the role in locals for potential use by other middlewares
	c.Locals("userRole", role)
	c.Locals("userID", userID)
	log.Printf("User %s role from JWT: %s", userID, role)

	if role != "admin" && role != "superadmin" {
		return c.Status(403).JSON(fiber.Map{"error": "Forbidden"})
	}
	return c.Next()
}

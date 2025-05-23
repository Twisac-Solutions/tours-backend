package routes

import (
	"github.com/Twisac-Solutions/tours-backend/controllers"
	"github.com/Twisac-Solutions/tours-backend/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	auth.Post("/google", controllers.GoogleSSO)

	user := api.Group("/user", middlewares.JWTProtected())
	user.Get("/profile", controllers.GetUserProfile)
}

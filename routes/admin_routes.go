package routes

import (
	"github.com/Twisac-Solutions/tours-backend/controllers"
	"github.com/Twisac-Solutions/tours-backend/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterAdminRoutes(app *fiber.App) {
	admin := app.Group("/admin", middlewares.AdminOnly)

	// Tour Routes
	admin.Get("/tours", controllers.GetAllTours)
	admin.Get("/tours/:id", controllers.GetTourByID)
	admin.Post("/tours", controllers.CreateTour)
	admin.Put("/tours/:id", controllers.UpdateTour)
	admin.Delete("/tours/:id", controllers.DeleteTour)
}

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

	//Events Routes
	admin.Get("/events", controllers.GetAllEvents)
	admin.Get("/events/:id", controllers.GetEventByID)
	admin.Post("/events", controllers.CreateEvent)
	admin.Put("/events/:id", controllers.UpdateEvent)
	admin.Delete("/events/:id", controllers.DeleteEvent)

	// Destination Routes
	admin.Get("/destinations", controllers.GetAllDestinations)
	admin.Get("/destinations/:id", controllers.GetDestinationByID)
	admin.Post("/destinations", controllers.CreateDestination)
	admin.Put("/destinations/:id", controllers.UpdateDestination)
	admin.Delete("/destinations/:id", controllers.DeleteDestination)

}

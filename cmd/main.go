package main

import (
	"log"

	"github.com/Twisac-Solutions/tours-backend/config"
	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/routes"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()

	app := fiber.New()
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}

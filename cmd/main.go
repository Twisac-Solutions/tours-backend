package main

import (
	"log"

	"github.com/Twisac-Solutions/tours-backend/config"
	"github.com/Twisac-Solutions/tours-backend/database"
	_ "github.com/Twisac-Solutions/tours-backend/docs"
	"github.com/Twisac-Solutions/tours-backend/routes"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Tours Backend API
// @version 1.0
// @description This is the API for the Tours Backend.
// @host localhost:8000
// @BasePath /
func main() {
	config.LoadEnv()
	database.ConnectDB()

	app := fiber.New()
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}

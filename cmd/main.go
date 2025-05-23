package main

import (
	"log"

	"github.com/Twisac-Solutions/tours-backend/config"
	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()

	app := fiber.New()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

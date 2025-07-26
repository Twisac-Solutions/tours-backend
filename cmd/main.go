package main

import (
	"log"

	"github.com/Twisac-Solutions/tours-backend/config"
	"github.com/Twisac-Solutions/tours-backend/database"

	// _ "github.com/Twisac-Solutions/tours-backend/docs"
	"github.com/Twisac-Solutions/tours-backend/routes"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/garrettladley/fiberpaginate/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Tours Backend API
// @version 1.0
// @description This is the API for the Tours Backend.
// @host localhost:8000
// @BasePath /
func main() {
	config.InitConfig()
	database.ConnectDB()
	database.SeedSuperAdmin()
	database.MigrateDB()

	err := utils.InitCloudinary()
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %v", err)
	}

	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024, // 10MB limit
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000,https://tours-dashboard-pi.vercel.app", // or your Next.js URL
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))
	// app.Get("/swagger/*", fiberSwagger.WrapHandler)
	app.Static("/docs", "./docs")
	app.Use(fiberpaginate.New())
	routes.SetupRoutes(app)
	routes.RegisterAdminRoutes(app)

	log.Fatal(app.Listen(":8000"))
}

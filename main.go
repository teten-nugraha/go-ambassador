package main

import (
	"ambassador/src/database"
	"ambassador/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	// Default middleware config
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":3000")
}

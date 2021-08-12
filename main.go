package main

import (
	"ambassador/src/database"
	"ambassador/src/routes"
	"ambassador/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()
	database.AutoMigrate()
	database.SetupRedis()

	app := fiber.New()
	utils.AddLogger(app)
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":3000")
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	_, err := gorm.Open(mysql.Open("root:root@tpc(db:3306)/ambassador"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}
	app := fiber.New()

	// Default middleware config
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}

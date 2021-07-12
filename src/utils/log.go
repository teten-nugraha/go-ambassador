package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func AddLogger(app *fiber.App) {

	//file, err := os.OpenFile("./app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalf("error opening file: %v", err)
	//}
	//defer file.Close()

	//app.Use(logger.New(logger.Config{
	//	Output: file,
	//}))

	app.Use(logger.New())
}

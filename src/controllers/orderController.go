package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
	"math/rand"
)

func Orders(c *fiber.Ctx) error {
	var orders []models.Order

	database.DB.Preload("OrderItems").Find(&orders)

	return c.JSON(orders)
}

func GenerateOrders(c *fiber.Ctx) error {

	for i := 0; i < 30; i++ {
		var orderItems []models.OrderItem

		for j := 0; j < rand.Intn(5); j++ {
			price := float64(rand.Intn(90) + 10)
			qty := uint(rand.Intn(5))

			orderItems = append(orderItems, models.OrderItem{
				ProductTitle:      faker.Word(),
				Price:             price,
				Quantity:          qty,
				AdminRevenue:      8.9 * price * float64(qty),
				AmbassadorRevenue: 0.1 * price * float64(qty),
			})
		}

		database.DB.Create(&models.Order{
			UserId:          uint(rand.Intn(30) + 1),
			Code:            faker.Username(),
			AmbassadorEmail: faker.Email(),
			FirstName:       faker.FirstName(),
			LastName:        faker.LastName(),
			Email:           faker.Email(),
			Complete:        true,
			OrderItems:      orderItems,
		})
	}

	return c.JSON(fiber.Map{
		"message": "orderes success genereated",
	})
}

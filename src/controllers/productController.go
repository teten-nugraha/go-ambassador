package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"strconv"
)

func Products(c *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)

	return c.JSON(products)
}

func CreateProducts(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var product models.Product

	product.Id = uint(id)

	database.DB.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id:          uint(id),
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(&product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id:          uint(id),
	}

	database.DB.Delete(&product)

	return c.JSON(fiber.Map{
		"message": "product success deleted",
	})
}

func PopulateProducts(c *fiber.Ctx) error {
	for i:=0; i < 30; i++ {
		product := models.Product{
			Title:       faker.Username(),
			Description: faker.Username(),
			Image:       faker.URL(),
			Price:       float64(rand.Intn(90)+10),
		}

		database.DB.Create(&product)
	}

	return c.JSON(fiber.Map{
		"message": "success generated products",
	})
}
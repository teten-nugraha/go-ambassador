package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
)

func Ambassadors(c *fiber.Ctx) error {
	var users[]models.User

	database.DB.Where("is_ambassador = true").Find(&users)

	return c.JSON(users)
}

func PopulateAmbassadors(c *fiber.Ctx) error {
	for i:=0; i < 30; i++ {
		ambassador := models.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}
		ambassador.SetPassword("1234")

		database.DB.Create(&ambassador)
	}

	return c.JSON(fiber.Map{
		"message": "success generated ambasadors users",
	})
}
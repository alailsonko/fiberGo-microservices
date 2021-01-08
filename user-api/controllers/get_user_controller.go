package controllers

import (
	"github.com/gofiber/fiber/v2"
	"user-api.com/database"
	"user-api.com/models"
)

func GETUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user models.User

	db.Find(&user, id)

	if int(user.ID) == 0 {
		c.JSON(fiber.Map{"message": "not found"})
		c.Status(404)
		return nil
	}

	c.JSON(fiber.Map{"data": user})
	c.Status(200)
	return nil
}

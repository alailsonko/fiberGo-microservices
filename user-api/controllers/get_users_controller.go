package controllers

import (
	"github.com/gofiber/fiber/v2"
	"user-api.com/database"
	"user-api.com/models"
)

func GETUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []models.User

	db.Find(&users)
	c.JSON(fiber.Map{"data": users})
	c.Status(200)
	return nil
}

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"order-api.com/database"
	"order-api.com/models"
)

func GETOrders(c *fiber.Ctx) error {
	db := database.DB
	var orders []models.Order

	db.Find(&orders)
	c.JSON(fiber.Map{"data": orders})
	c.Status(200)
	return nil
}

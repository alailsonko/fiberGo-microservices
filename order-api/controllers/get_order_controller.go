package controllers

import (
	"github.com/gofiber/fiber/v2"
	"order-api.com/database"
	"order-api.com/models"
)

func GETOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var order models.Order

	db.Find(&order, id)

	if int(order.ID) == 0 {
		c.JSON(fiber.Map{"message": "not found"})
		c.Status(404)
		return nil
	}

	c.JSON(fiber.Map{"data": order})
	c.Status(200)
	return nil
}

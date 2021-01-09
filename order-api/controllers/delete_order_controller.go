package controllers

import (
	"github.com/gofiber/fiber/v2"
	"order-api.com/database"
	"order-api.com/models"
)

func DELETEOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var order models.Order

	db.First(&order, id)

	if order.ID == 0 {
		c.JSON(fiber.Map{"message": "not found"})
		c.Status(404)
		return nil
	}

	db.Delete(&order)

	c.JSON(fiber.Map{"message": "deleted successfully"})
	c.Status(200)

	return nil
}

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"order-api.com/cache"
)

func GETOrders(c *fiber.Ctx) error {
	var cc *cache.Cache

	orders, err := cc.GetOrdersCache()

	if err != nil {
		c.JSON(fiber.Map{"error": err})
		c.Status(fiber.StatusBadRequest)
		return nil
	}

	c.JSON(fiber.Map{"data": orders})
	c.Status(200)
	return nil
}

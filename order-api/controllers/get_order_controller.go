package controllers

import (
	"github.com/gofiber/fiber/v2"
	"order-api.com/cache"
)

func GETOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var cc *cache.Cache

	order, err := cc.GetOrderCache(id)
	if err != nil {
		c.JSON(fiber.Map{"err": err})
		c.Status(fiber.StatusBadRequest)
		return nil
	}
	c.JSON(fiber.Map{"data": order})
	c.Status(200)
	return nil
}

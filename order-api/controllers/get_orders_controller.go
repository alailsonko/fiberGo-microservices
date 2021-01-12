package controllers

import (
	"github.com/gofiber/fiber/v2"
	"order-api.com/cache"
)

// GETOrders godoc
// @Summary Get orders
// @Description get all orders
// @Produce  json
// @Success 200 {object} []models.Order
// @Failure 404 {object} docs.HTTPError
// @Router /order [get]
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

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"order-api.com/cache"
)

// GETOrder godoc
// @Summary get an order
// @Description get order by ID
// @ID get-order-by-int
// @Produce  json
// @Param id path string true "Order ID"
// @Success 200 {object} models.Order
// @Failure 404 {object} docs.HTTPError
// @Router /order/:id [get]
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

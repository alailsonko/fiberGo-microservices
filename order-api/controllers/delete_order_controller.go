package controllers

import (
	"github.com/gofiber/fiber/v2"
	"order-api.com/cache"
	"order-api.com/database"
	"order-api.com/models"
)

// DELETETOrder godoc
// @Summary delete an order
// @Description delete order by ID
// @ID delete-order-by-string
// @Produce  json
// @Param id path string true "OrderID"
// @Success 200 {object} docs.HTTPOk
// @Failure 404 {object} docs.HTTPError
// @Router /order/:id [delete]
func DELETEOrder(c *fiber.Ctx) error {
	var cc *cache.Cache
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

	cc.UpdateOrdersCache()
	c.JSON(fiber.Map{"message": "deleted successfully"})
	c.Status(200)
	return nil
}

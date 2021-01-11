package controllers

import (
	"github.com/gofiber/fiber/v2"
	"order-api.com/cache"
	"order-api.com/database"
	"order-api.com/utils"
	"order-api.com/validators"
)

// Order - types
type Order struct {
	*validators.Order
}

// CREATEOrder - handler for create order
func CREATEOrder(c *fiber.Ctx) error {
	var cc *cache.Cache
	db := database.DB

	o := new(Order)

	if err := c.BodyParser(o); err != nil {
		return err
	}

	err := o.ValidateOrder()

	if err != nil {
		c.JSON(fiber.Map{"error": err.Error()})
		c.SendStatus(400)
		return nil
	}

	o.TotalPrice = utils.TotalPrice(o.ItemQuantity, o.ItemPrice)
	dd := db.Create(&o)

	if dd.Error != nil {
		c.JSON(fiber.Map{"error": dd.Error.Error()})
		c.SendStatus(400)
		return nil
	}
	cc.UpdateOrdersCache()
	c.JSON(fiber.Map{"msg": "Order created successfully."})
	c.SendStatus(200)
	return nil
}

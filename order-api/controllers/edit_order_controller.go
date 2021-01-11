package controllers

import (
	"github.com/gofiber/fiber/v2"
	"order-api.com/cache"
	"order-api.com/database"
	"order-api.com/models"
	"order-api.com/utils"
)

// EDITOrder - handler for create order
func EDITOrder(c *fiber.Ctx) error {
	var cc *cache.Cache
	id := c.Params("id")
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
	dd := db.Model(&models.Order{}).Where("id = ?", id).Updates(models.Order{UserID: o.UserID, ItemDescription: o.ItemDescription, ItemQuantity: o.ItemQuantity, ItemPrice: o.ItemPrice, TotalPrice: o.TotalPrice})

	if dd.Error != nil {
		c.JSON(fiber.Map{"error": dd.Error.Error()})
		c.SendStatus(400)
		return nil
	}
	cc.UpdateOrdersCache()
	c.JSON(fiber.Map{"msg": "User edited successfully."})
	c.SendStatus(200)
	return nil
}

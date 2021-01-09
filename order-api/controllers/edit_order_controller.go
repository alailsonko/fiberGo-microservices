package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"order-api.com/database"
	"order-api.com/models"
	"order-api.com/utils"
)

// EDITOrder - handler for create order
func EDITOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	o := new(Order)

	if err := c.BodyParser(o); err != nil {
		return err
	}

	err := o.ValidateOrder()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(fiber.Map{"error": err.Error()})
		c.SendStatus(400)
		return nil
	}

	log.Println("works")
	o.TotalPrice = utils.TotalPrice(o.ItemQuantity, o.ItemPrice)
	dd := db.Model(&models.Order{}).Where("id = ?", id).Updates(models.Order{UserID: o.UserID, ItemDescription: o.ItemDescription, ItemQuantity: o.ItemQuantity, ItemPrice: o.ItemPrice, TotalPrice: o.TotalPrice})
	log.Println("works", dd.Error)

	if dd.Error != nil {
		fmt.Println(dd.Error)
		c.JSON(fiber.Map{"error": dd.Error.Error()})
		c.SendStatus(400)
		return nil
	}

	c.JSON(fiber.Map{"msg": "User edited successfully."})
	log.Println("works")

	c.SendStatus(200)
	log.Println("works")

	return nil
}

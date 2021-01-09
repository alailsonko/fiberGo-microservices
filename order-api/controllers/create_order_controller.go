package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
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
	db := database.DB

	o := new(Order)

	log.Println("o", (*o))

	if err := c.BodyParser(o); err != nil {
		log.Println("o", o)
		log.Println("err", err)
		return err
	}

	err := o.ValidateOrder()

	log.Println(err)

	if err != nil {
		c.JSON(fiber.Map{"error": err.Error()})
		c.SendStatus(400)
		return nil
	}

	log.Println(o.UserID)
	log.Println(o.ItemDescription)
	log.Println(o.ItemPrice)
	log.Println(o.ItemQuantity)
	o.TotalPrice = utils.TotalPrice(o.ItemQuantity, o.ItemPrice)
	log.Println("total", o.TotalPrice)
	dd := db.Create(&o)
	log.Println("dd", dd)

	if dd.Error != nil {
		log.Println("dd.Error", dd.Error)

		c.JSON(fiber.Map{"error": dd.Error.Error()})
		c.SendStatus(400)
		return nil
	}

	c.JSON(fiber.Map{"msg": "Order created successfully."})
	c.SendStatus(200)
	return nil
}

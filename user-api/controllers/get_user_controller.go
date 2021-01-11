package controllers

import (
	"github.com/gofiber/fiber/v2"
	"user-api.com/cache"
)

// GETUser controller
func GETUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var d *cache.Cache
	user, err := d.GetUserCache(id)

	if err != nil {
		c.JSON(fiber.Map{"error": err.Error()})
		c.Status(404)
		return nil
	}

	c.JSON(fiber.Map{"data": user})
	c.Status(200)
	return nil
}

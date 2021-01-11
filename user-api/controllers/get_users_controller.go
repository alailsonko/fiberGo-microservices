package controllers

import (
	"github.com/gofiber/fiber/v2"
	"user-api.com/cache"
)

func GETUsers(c *fiber.Ctx) error {
	var d *cache.Cache
	users, err := d.GetUsersCache()

	if err != nil {
		c.JSON(fiber.Map{"error": err})
		c.Status(404)
		return nil
	}
	c.JSON(fiber.Map{"data": users})
	c.Status(200)
	return nil
}

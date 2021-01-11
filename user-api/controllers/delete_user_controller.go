package controllers

import (
	"github.com/gofiber/fiber/v2"
	"user-api.com/cache"
	"user-api.com/database"
	"user-api.com/models"
)

func DELETEUser(c *fiber.Ctx) error {
	var cc *cache.Cache

	id := c.Params("id")
	db := database.DB

	var user models.User

	db.First(&user, id)

	if user.ID == 0 {
		c.JSON(fiber.Map{"message": "not found"})
		c.Status(404)
		return nil
	}
	db.Delete(&user)

	c.JSON(fiber.Map{"message": "delete successfully"})
	c.Status(200)
	cc.UpdateUsersCache()
	return nil
}

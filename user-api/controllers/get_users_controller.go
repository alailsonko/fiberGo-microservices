package controllers

import (
	"github.com/gofiber/fiber/v2"
	"user-api.com/cache"
)

// GETUsers godoc
// @Summary Get users
// @Description get all users
// @Produce  json
// @Success 200 {object} []models.User
// @Failure 404 {object} docs.HTTPError
// @Router /user [get]
func GETUsers(c *fiber.Ctx) error {
	var cc *cache.Cache
	users, err := cc.GetUsersCache()

	if err != nil {
		c.JSON(fiber.Map{"error": err})
		c.Status(fiber.StatusBadRequest)
		return nil
	}
	c.JSON(fiber.Map{"data": users})
	c.Status(fiber.StatusOK)
	return nil
}

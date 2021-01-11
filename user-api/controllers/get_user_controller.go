package controllers

import (
	"github.com/gofiber/fiber/v2"
	"user-api.com/cache"
)

// GETUser godoc
// @Summary get an user
// @Description get user by ID
// @ID get-user-by-int
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} docs.HTTPError
// @Router /user/{id} [get]
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

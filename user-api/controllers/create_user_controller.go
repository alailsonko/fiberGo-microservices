package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"user-api.com/cache"
	"user-api.com/database"
	"user-api.com/validators"
)

// User - types of the user
type User struct {
	*validators.User
}

// CREATEUser godoc
// @Summary create an user
// @Description create user by ID
// @ID create-user-by-string
// @Accept   json
// @Produce  json
// @Param user body docs.User true "accepts user data"
// @Success 200 {object} docs.HTTPOk
// @Failure 404 {object} docs.HTTPError
// @Router /user [post]
func CREATEUser(c *fiber.Ctx) error {
	var cc *cache.Cache

	db := database.DB

	u := new(User)

	if err := c.BodyParser(u); err != nil {
		return err
	}

	err := u.ValidateUser()
	if err != nil {
		c.JSON(fiber.Map{"error": err.Error()})
		c.SendStatus(400)
		return nil
	}

	dd := db.Create(&u)
	log.Println("works", dd.Error)

	if dd.Error != nil {
		c.JSON(fiber.Map{"error": dd.Error.Error()})
		c.SendStatus(400)
		return nil
	}
	cc.UpdateUsersCache()
	c.JSON(fiber.Map{"msg": "User created successfully."})
	c.SendStatus(200)
	return nil
}

package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"user-api.com/cache"
	"user-api.com/database"
	"user-api.com/models"
)

// EDITUser godoc
// @Summary edit an user
// @Description edit user by ID
// @ID edit-user-by-string
// @Accept   json
// @Produce  json
// @Param id path string true "UserID"
// @Param user body docs.User true "accepts user data"
// @Success 200 {object} docs.HTTPOk
// @Failure 404 {object} docs.HTTPError
// @Router /user/{id} [put]
func EDITUser(c *fiber.Ctx) error {
	var cc *cache.Cache

	id := c.Params("id")
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

	dd := db.Model(&models.User{}).Where("id = ?", id).Updates(models.User{PhoneNumber: u.PhoneNumber, Cpf: u.Cpf, Email: u.Email, Username: u.Username})
	log.Println("works", dd.Error)

	if dd.Error != nil {
		fmt.Println(dd.Error)
		c.JSON(fiber.Map{"error": dd.Error.Error()})
		c.SendStatus(400)
		return nil
	}
	cc.UpdateUsersCache()
	c.JSON(fiber.Map{"msg": "User edited successfully."})
	c.SendStatus(200)
	return nil
}

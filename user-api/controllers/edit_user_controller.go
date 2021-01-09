package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"user-api.com/database"
	"user-api.com/models"
)

// // User - types of the user
// type User struct {
// 	*validators.User
// }

// EDITUser - handler for create user
func EDITUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	u := new(User)

	if err := c.BodyParser(u); err != nil {
		return err
	}

	log.Println(u.Username)
	log.Println(u.Cpf)
	log.Println(u.Email)
	log.Println(u.PhoneNumber)

	err := u.ValidateUser()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(fiber.Map{"error": err.Error()})
		c.SendStatus(400)
		return nil
	}

	log.Println("works")

	dd := db.Model(&models.User{}).Where("id = ?", id).Updates(models.User{PhoneNumber: u.PhoneNumber, Cpf: u.Cpf, Email: u.Email, Username: u.Username})
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

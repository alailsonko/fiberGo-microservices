package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"user-api.com/validators"
)

// User - types of the user
type User struct {
	validators.User
}

// CREATEUser - handler for create user
func CREATEUser(c *fiber.Ctx) error {
	// db := database.DB
	ut := new(User)

	if err := c.BodyParser(ut); err != nil {
		return err
	}

	log.Println(ut.Username)
	log.Println(ut.Cpf)
	log.Println(ut.Email)
	log.Println(ut.PhoneNumber)

	err := ut.ValidateUser()

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(fiber.Map{"error": err.Error()})
		c.SendStatus(400)
		return nil
	}

	c.JSON(fiber.Map{"msg": "User created successfully."})
	c.SendStatus(400)
	return nil
}

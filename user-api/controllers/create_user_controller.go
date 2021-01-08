package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"user-api.com/database"
	"user-api.com/validators"
)

// User - types of the user
type User struct {
	*validators.User
}

// CREATEUser - handler for create user
func CREATEUser(c *fiber.Ctx) error {
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

	dd := db.Create(&u)
	log.Println("works", dd)

	c.JSON(fiber.Map{"msg": "User created successfully."})
	log.Println("works")

	c.SendStatus(200)
	log.Println("works")

	return nil
}

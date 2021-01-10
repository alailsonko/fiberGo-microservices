package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"order-api.com/database"
	"order-api.com/models"
)

// User - response data of the user
type User struct {
	gorm.Model
	Username    string `json:"username"`
	Email       string `json:"email"`
	Cpf         string `json:"cpf"`
	PhoneNumber string `json:"phone_number"`
}

// Response - of the user service
type Response struct {
	Data User `json:"data"`
}

// GETOrdersByUser - get all orders by user
func GETOrdersByUser(c *fiber.Ctx) error {

	idUser := c.Params("id")

	db := database.DB

	response, err := http.Get("http://172.21.0.1:3030/user/" + idUser)
	if err != nil {
		c.JSON(fiber.Map{"data": "not found"})
		c.Status(404)
		return nil
	}

	user, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(fiber.Map{"data": "not found"})
		c.Status(404)
		return nil
	}

	var jsonMapUser Response
	json.Unmarshal(user, &jsonMapUser)
	json.Marshal(jsonMapUser)

	var orders []models.Order

	log.Println(idUser)

	db.Raw("SELECT * FROM orders WHERE user_id = ?", idUser).Scan(&orders)

	if int(len(orders)) == 0 {
		c.JSON(fiber.Map{"message": "not found"})
		c.Status(404)
		return nil
	}

	c.JSON(fiber.Map{"user": jsonMapUser.Data, "orders": orders})
	c.Status(200)
	return nil
}

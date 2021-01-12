package repositories

import (
	"errors"
	"log"

	"order-api.com/database"
	"order-api.com/models"
)

// GETOrderRepository using param id
func GETOrderRepository(id string) (*models.Order, error) {

	db := database.DB
	order := &models.Order{}

	dd := db.Find(&order, id)
	log.Println("dd", dd)

	if int(order.ID) == 0 {
		err := errors.New("not found")
		log.Println("err", err)
		return order, err
	}

	return order, dd.Error
}

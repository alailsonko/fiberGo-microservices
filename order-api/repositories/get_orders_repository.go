package repositories

import (
	"order-api.com/database"
	"order-api.com/models"
)

func GETOrdersRepository() ([]models.Order, error) {
	db := database.DB
	var orders []models.Order

	dd := db.Find(&orders)

	return orders, dd.Error
}

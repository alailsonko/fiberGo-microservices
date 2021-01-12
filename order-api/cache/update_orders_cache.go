package cache

import (
	"encoding/json"

	"order-api.com/models"
	"order-api.com/repositories"
)

// UpdateOrdersCache when value is not found
func (c *Cache) UpdateOrdersCache() ([]models.Order, error) {
	data, err := repositories.GETOrdersRepository()

	if err != nil {

		return nil, err
	}
	byteData, _ := json.Marshal(&data)

	if err := c.Set("orders", byteData); err != nil {

		return nil, err
	}

	return data, nil
}

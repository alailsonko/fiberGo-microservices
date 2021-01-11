package cache

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/go-redis/redis"
	"order-api.com/models"
	"order-api.com/repositories"
)

// GetOrderCache by id
func (c *Cache) GetOrderCache(id string) (*models.Order, error) {

	b, err := c.Get("orders")

	if err != nil && err != redis.Nil {
		return nil, err
	}

	if err == redis.Nil {
		c.UpdateOrdersCache()
		data, err := repositories.GETOrderRepository(id)

		return data, err
	}

	var data []models.Order
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	// Build a config map:
	for _, u := range data {
		strID := strconv.FormatUint(uint64(u.ID), 10)
		if strID == id {
			return &u, nil
		}
	}
	err = errors.New("not found")
	return nil, err

}

package cache

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"order-api.com/models"
)

// GetOrdersCache implements logic control of caching
func (c *Cache) GetOrdersCache() ([]models.Order, error) {
	b, err := c.Get("orders")

	if err != nil && err != redis.Nil {
		return nil, err
	}

	if err == redis.Nil {
		data, err := c.UpdateOrdersCache()
		return data, err
	}

	var data []models.Order

	if err := json.Unmarshal(b, &data); err != nil {

		return nil, err
	}

	return data, nil
}

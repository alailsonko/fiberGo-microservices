package cache

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"user-api.com/models"
)

// GetUsersCache implements logic control of caching
func (c *Cache) GetUsersCache() ([]models.User, error) {
	b, err := c.Get("users")

	if err != nil && err != redis.Nil {
		return nil, err
	}

	if err == redis.Nil {
		data, err := c.UpdateUsersCache()
		return data, err
	}

	var data []models.User

	if err := json.Unmarshal(b, &data); err != nil {

		return nil, err
	}

	return data, nil
}

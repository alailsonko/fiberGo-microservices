package cache

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/go-redis/redis"
	"user-api.com/models"
	"user-api.com/repositories"
)

// GetUserCache by id
func (c *Cache) GetUserCache(id string) (*models.User, error) {

	b, err := c.Get("users")

	if err != nil && err != redis.Nil {
		return nil, err
	}

	if err == redis.Nil {
		c.UpdateUsersCache()
		data, err := repositories.GETUserRepository(id)

		return data, err
	}

	var data []models.User
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

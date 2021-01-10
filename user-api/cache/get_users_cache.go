package cache

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"user-api.com/models"
	"user-api.com/repositories"
)

func (c *Cache) GetUsersCache() ([]models.User, error) {
	b, err := c.Get("users")

	if err != nil && err != redis.Nil {
		return nil, err
	}

	if err == redis.Nil {

		data, err := repositories.GETUsersRepository()

		if err != nil {

			return nil, err
		}
		byteData, _ := json.Marshal(&data)
		if err := c.Set("users", byteData); err != nil {

			return nil, err
		}

		return data, nil
	}

	var data []models.User

	if err := json.Unmarshal(b, &data); err != nil {

		return nil, err
	}

	return data, nil
}

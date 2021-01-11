package cache

import (
	"encoding/json"

	"user-api.com/models"
	"user-api.com/repositories"
)

// UpdateUsersCache when value is not found
func (c *Cache) UpdateUsersCache() ([]models.User, error) {
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

package cache

import (
	"github.com/go-redis/redis"
)

// Interface create interface that should be respected
type Interface interface {
	Get(key string) ([]byte, error)
	Set(key string, value interface{}) error
}

// Cache set methods
type Cache struct {
	Cache *redis.Client
}

// RClient connect to redis
func (c *Cache) RClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "0799dffe70a8:6379",
		Password: "",
		DB:       0,
	})

	return client
}

// Get values by key
func (c *Cache) Get(key string) ([]byte, error) {
	s, err := c.RClient().Get(key).Result()

	if err != nil {
		return nil, err
	}
	data := []byte(s)

	return data, nil
}

// Set keys and values
func (c *Cache) Set(key string, value interface{}) error {
	err := c.RClient().Set(key, value, 0).Err()

	if err != nil {
		return err
	}

	return nil
}

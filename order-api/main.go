package main

import (
	"fmt"
	"os"

	redis "gopkg.in/redis.v4"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

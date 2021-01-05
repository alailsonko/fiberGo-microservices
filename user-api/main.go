package main

import (
	"log"
	"os"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello")
	})

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

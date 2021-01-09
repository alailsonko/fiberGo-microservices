package setup_app

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupApp() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())

	return app
}

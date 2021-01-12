package setup_app

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupApp() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(cors.New())
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://localhost:3031/swagger/doc.json",
		DeepLinking: false,
	}))
	return app
}

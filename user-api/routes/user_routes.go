package routes

import (
	"github.com/gofiber/fiber/v2"
	"user-api.com/controllers"
)

// UserSetupRoutes - sets all routes for user
func UserSetupRoutes(app *fiber.App) {
	app.Post("/user", controllers.CREATEUser)
	app.Get("/user", controllers.GETUsers)
	app.Get("/user/:id", controllers.GETUser)
	app.Delete("/user/:id", controllers.DELETEUser)
}

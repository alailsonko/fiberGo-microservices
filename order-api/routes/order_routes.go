package routes

import (
	"github.com/gofiber/fiber/v2"
	"order-api.com/controllers"
)

// OrderSetupRoutes - sets all routes for order
func OrderSetupRoutes(app *fiber.App) {
	app.Post("/order", controllers.CREATEOrder)
	app.Get("/order/:id", controllers.GETOrder)
	app.Delete("/order/:id", controllers.DELETEOrder)
	app.Put("/order/:id", controllers.EDITOrder)
	app.Get("/orders-by-user/:id", controllers.GETOrdersByUser)
	app.Get("/order", controllers.GETOrders)
}

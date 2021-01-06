package main

import (
	"log"
	"os"

	"user-api.com/routes"
	"user-api.com/setup_app"
)

func main() {
	app := setup_app.SetupApp()

	routes.UserSetupRoutes(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

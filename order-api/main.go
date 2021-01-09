package main

import (
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"order-api.com/database"
	"order-api.com/routes"
	"order-api.com/setup_app"
)

func main() {

	app := setup_app.SetupApp()
	database.ConnectPGDB()
	routes.OrderSetupRoutes(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

package main

import (
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"user-api.com/database"
	"user-api.com/routes"
	"user-api.com/setup_app"
)

func main() {

	app := setup_app.SetupApp()
	database.ConnectPGDB()
	routes.UserSetupRoutes(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

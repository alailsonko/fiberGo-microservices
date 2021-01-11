package main

import (
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"user-api.com/database"
	_ "user-api.com/docs"
	"user-api.com/routes"
	"user-api.com/setup_app"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3030
// @BasePath /
func main() {

	app := setup_app.SetupApp()
	database.ConnectPGDB()
	routes.UserSetupRoutes(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

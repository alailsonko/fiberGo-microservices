package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	"user-api.com/config"
	"user-api.com/models"
)

// DB gorm connector
var DB *gorm.DB

// ConnectPGDB connect to db
func ConnectPGDB() {
	var err error

	p := config.Config("POSTGRES_PORT")

	port, err := strconv.ParseUint(p, 10, 32)

	log.Printf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("POSTGRES_HOST"), port, config.Config("POSTGRES_USER"), config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_DB"))

	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("POSTGRES_HOST"), port, config.Config("POSTGRES_USER"), config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_DB")))

	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connection Opened to Database")
	DB.AutoMigrate(&models.User{})
	log.Println("Database Migrated")
}

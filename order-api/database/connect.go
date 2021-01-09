package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	"order-api.com/config"
	"order-api.com/models"
)

// DB gorm connector
var DB *gorm.DB

// ConnectPGDB connect to db
func ConnectPGDB() {
	var err error

	p := config.Config("POSTGRES_PORT")

	log.Println("p", p)

	port, err := strconv.ParseUint(p, 10, 32)
	log.Println("port", port)

	log.Println("err", err)

	log.Printf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("POSTGRES_HOST"), port, config.Config("POSTGRES_USER"), config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_DB"))
	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("POSTGRES_HOST"), port, config.Config("POSTGRES_USER"), config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_DB")))
	log.Println("DB", DB)

	log.Println("err", err)

	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connection Opened to Database")
	dd := DB.AutoMigrate(&models.Order{})
	log.Println("Database Migrated", dd)
}

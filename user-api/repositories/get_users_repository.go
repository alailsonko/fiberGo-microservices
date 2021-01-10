package repositories

import (
	"user-api.com/database"
	"user-api.com/models"
)

// GETUsersRepository - gets all users
func GETUsersRepository() ([]models.User, error) {
	db := database.DB
	var users []models.User

	dd := db.Find(&users)

	return users, dd.Error
}

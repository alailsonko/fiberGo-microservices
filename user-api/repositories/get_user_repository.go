package repositories

import (
	"errors"
	"log"

	"user-api.com/database"
	"user-api.com/models"
)

// GETUserRepository using param id
func GETUserRepository(id string) (*models.User, error) {

	db := database.DB
	user := &models.User{}

	dd := db.Find(&user, id)
	log.Println("dd", dd)

	if int(user.ID) == 0 {
		err := errors.New("not found")
		log.Println("err", err)
		return user, err
	}

	return user, dd.Error
}

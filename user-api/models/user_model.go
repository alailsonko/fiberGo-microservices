package models

import (
	"github.com/jinzhu/gorm"
)

// UserModel - struct
type User struct {
	gorm.Model

	Username    string `gorm:"not null" json:"username"`
	Cpf         string `gorm:"not null" json:"cpf"`
	Email       string `gorm:"not null" json:"email"`
	PhoneNumber string `gorm:"not null" json:"phone_number"`
}

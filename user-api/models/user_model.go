package models

import (
	"github.com/jinzhu/gorm"
)

// User - struct
type User struct {
	gorm.Model

	Username    string `gorm:"not null" json:"username"`
	Cpf         string `gorm:"not null" gorm:"unique" json:"cpf"`
	Email       string `gorm:"not null" gorm:"unique" json:"email"`
	PhoneNumber string `gorm:"not null" json:"phone_number"`
}

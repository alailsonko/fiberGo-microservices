package models

import (
	"github.com/jinzhu/gorm"
)

// Order - struct
type Order struct {
	gorm.Model

	UserID          string  `gorm:"not null" json:"user_id"`
	ItemDescription string  `gorm:"not null" json:"item_description"`
	ItemQuantity    float64 `gorm:"not null" json:"item_item_quantity"`
	ItemPrice       float64 `gorm:"not null" json:"item_price"`
	TotalPrice      float64 `gorm:"not null" json:"total_price"`
}

package models

import "time"

type Score struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Rate      uint      `json:"rate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ProductID uint      `json:"productId"`
	UserID    uint      `json:"userId"`
}

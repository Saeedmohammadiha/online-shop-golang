package models

import "time"

type Discount struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Amount    float64   `json:"amount" gorm:"type:decimal(19,4);"`
	Percent   float64   `json:"percent" gorm:"type:decimal(5,4);"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ProductID uint      `json:"ProductId"`
	UserID    uint      `json:"userId"`
}

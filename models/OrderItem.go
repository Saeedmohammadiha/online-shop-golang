package models

import "time"

type OrderItem struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Quantity  uint      `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updated_at"`
	ProductID uint      `json:"productId"`
	OrderID   uint      `json:"orderId"`
}

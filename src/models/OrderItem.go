package models

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	Quantity  uint `json:"quantity"`
	ProductID uint `json:"productId"`
	OrderID   uint `json:"orderId"`
}

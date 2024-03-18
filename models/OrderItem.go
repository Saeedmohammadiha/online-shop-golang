package models

import "time"

type OrderItem struct {
    ID        uint `gorm:"primaryKey"`
    ProductID uint
    OrderID   uint
    Quantity  int
    CreatedAt time.Time
}
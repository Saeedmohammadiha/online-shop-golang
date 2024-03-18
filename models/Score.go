package models

import "time"

type Score struct {
    ID        uint `gorm:"primaryKey"`
    ProductID uint
    UserID    uint
    Rate      int
    CreatedAt time.Time
    UpdatedAt time.Time
}
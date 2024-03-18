package models

import (
	"time"
)


type Product struct {
    ID          uint           `gorm:"primaryKey"`
    Name        string
    Price       float64
    DiscountID  uint
    Quantity    int
    CreatedAt   time.Time
    UpdatedAt   time.Time
    Summary     string         `gorm:"type:text"`
    Slug        string
}
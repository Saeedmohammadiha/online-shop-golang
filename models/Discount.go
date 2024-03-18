package models

import "time"

type Discount struct {
    ID        uint           `gorm:"primaryKey"`
    Amount    float64
    CreatedAt time.Time
    UpdatedAt time.Time
}
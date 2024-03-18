package models

import "time"

type Comment struct {
    ID        uint `gorm:"primaryKey"`
    Content   string         `gorm:"type:text"`
    ProductID uint
    UserID    uint
    CreatedAt time.Time
    UpdatedAt time.Time
}
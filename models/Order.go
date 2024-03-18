package models

import (
	"time"
)

type Order struct {
    ID                uint           `gorm:"primaryKey"`
    UserID            uint
    Tax               float64
    StatusID          uint
    AddressID         uint
    PostalTrackingCode int
    CreatedAt         time.Time
}
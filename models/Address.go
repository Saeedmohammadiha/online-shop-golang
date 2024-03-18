package models

import (
	"time"
)

type Address struct {
    ID          uint           `gorm:"primaryKey"`
    UserID      uint
    Province    string
    City        string
    Street      string
    PostalCode  int
    Description string         `gorm:"type:text"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    Users       []User        `gorm:"many2many:user_addresses;"`
}
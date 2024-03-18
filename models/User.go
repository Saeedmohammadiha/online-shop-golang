package models

import (
	"time"
)

type User struct {
    ID           uint           `gorm:"primaryKey"`
    Name         string
    LastName     string
    Email        string         `gorm:"unique"`
    PhoneNumber  string
    RoleID       uint
    DiscountID   uint
    PasswordHash string
    ActivatedAt  time.Time
    CreatedAt    time.Time
    UpdatedAt    time.Time
    Roles        []Role        `gorm:"many2many:user_roles;"`
    Addresses    []Address     `gorm:"many2many:user_addresses;"`
}
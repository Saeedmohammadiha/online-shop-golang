package models

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `json:"name"`
	LastName     string    `json:"lastName"`
	Email        string    `gorm:"unique" json:"email"`
	PhoneNumber  string    `json:"phoneNumber"`
	DiscountID   uint      `json:"discountId"`
	PasswordHash string    `json:"password"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Roles        []Role    `gorm:"many2many:user_roles;" json:"roles"`
	Addresses    []Address `gorm:"many2many:user_addresses;" json:"adresses"`
}

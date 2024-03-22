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
	DiscountID   *uint      `json:"discountId" gorm:"default:1"`
	PasswordHash string    `json:"password"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Roles        []Role    `json:"roles" gorm:"many2many:user_roles; default:1"`
	Addresses    []Address `gorm:"many2many:user_addresses;" json:"adresses"`
}

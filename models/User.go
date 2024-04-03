package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string        `json:"name"`
	LastName     string        `json:"lastName"`
	Email        string        `gorm:"unique" json:"email"`
	Password     string        `json:"password"`
	PhoneNumber  string        `json:"phoneNumber"`
	RoleID       int           `json:"roleId"`
	Transactions []Transaction `json:"transactions,omitempty"`
	Addresses    []Address     `gorm:"many2many:user_addresses" json:"addresses,omitempty"`
	Orders       []Order       `json:"orders,omitempty"`
	Scores       []Score       `json:"scores,omitempty"`
	Comments     []Comment     `json:"comments,omitempty"`
}

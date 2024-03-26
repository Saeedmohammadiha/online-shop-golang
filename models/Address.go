package models

import (
	"time"
)

type Address struct {
	ID          uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	Street      string    `json:"street"`
	PostalCode  uint      `json:"postalCode"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

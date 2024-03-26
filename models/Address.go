package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Province    string `json:"province"`
	City        string `json:"city"`
	Street      string `json:"street"`
	PostalCode  uint   `json:"postalCode"`
	Description string `json:"description"`
}

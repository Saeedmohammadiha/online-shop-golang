package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID      uint
	Province    string
	City        string
	Street      string
	PostalCode  string
	Description string
}

package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title       string
	Price       float64
	Pictures    string
	Description string
	Category    string 
}

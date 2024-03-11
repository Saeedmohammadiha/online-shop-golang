package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint
	Product []Product  `gorm:"many2many:order_products;"`
	Count int
}

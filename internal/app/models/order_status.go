package models

import "gorm.io/gorm"

type OrderStatus struct {
	gorm.Model
	Title string `json:"title"`
}

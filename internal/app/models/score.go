package models

import (
	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
	Rate      uint `json:"rate"`
	ProductID uint `json:"productId"`
	UserID    uint `json:"userId"`
}

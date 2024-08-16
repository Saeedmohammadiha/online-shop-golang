package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content   string `json:"content"`
	ProductID uint   `json:"productId"`
	UserID    uint   `json:"userId"`
}

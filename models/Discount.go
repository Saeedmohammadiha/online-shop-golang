package models

import "gorm.io/gorm"

type Discount struct {
	gorm.Model
	Amount    float64 `json:"amount" gorm:"type:decimal(19,4);"`
	Percent   float64 `json:"percent" gorm:"type:decimal(5,4);"`
	Code      string  `json:"code"`
	ProductID uint    `json:"ProductId"`
	UserID    uint    `json:"userId"`
}

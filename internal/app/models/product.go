package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name       string      `json:"name"`
	Price      float64     `json:"price"`
	Quantity   uint        `json:"quantity"`
	Summery    string      `json:"summery"`
	Slug       string      `json:"slug"`
	OrderItems []OrderItem `json:"orderItems,omitempty"`
	Scores     []Score     `json:"scores,omitempty"`
	Comments   []Comment   `json:"comments,omitempty"`
}

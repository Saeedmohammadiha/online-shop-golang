package models

import (
	"time"
)

type Product struct {
	ID         uint        `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name       string      `json:"name"`
	Price      float64     `json:"price"`
	Quantity   uint        `json:"quantity"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
	Summery    string      `json:"summery"`
	Slug       string      `json:"slug"`
	OrderItems []OrderItem `json:"orderItems,omitempty"` 
	Scores     []Score     `json:"scores,omitempty"`     
	Comments   []Comment   `json:"comments,omitempty"`   
}

package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ProductID uint      `json:"productId"`
	UserID    uint      `json:"userId"`
	//	Product   Product   `json:"product,omitempty"` // Optional relationship with Product model
	//	User      User      `json:"user,omitempty"`       // Optional relationship with User model
}

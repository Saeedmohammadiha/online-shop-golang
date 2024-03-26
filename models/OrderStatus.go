package models

type OrderStatus struct {
	ID    uint    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title string  `json:"title"`
}

package models

type Permission struct {
	ID    uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title string `json:"title"`
}

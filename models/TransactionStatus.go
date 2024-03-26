package models

type TransactionStatus struct {
	ID    uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title string `json:"title"`
}

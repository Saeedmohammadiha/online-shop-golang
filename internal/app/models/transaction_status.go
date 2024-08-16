package models

import "gorm.io/gorm"

type TransactionStatus struct {
	gorm.Model
	Title string `json:"title"`
}

package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Title string `json:"title"`
}

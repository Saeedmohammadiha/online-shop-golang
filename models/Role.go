package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Title       string       `json:"title"`
	Permissions []Permission `gorm:"many2many:permission_roles" json:"permissions,omitempty"` // Many-to-Many relationship with Permission model
}

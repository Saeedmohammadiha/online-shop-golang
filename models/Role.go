package models

type Role struct {
	ID          uint         `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title       string       `json:"title"`
	Permissions []Permission `gorm:"many2many:permission_roles" json:"permissions,omitempty"` // Many-to-Many relationship with Permission model
}

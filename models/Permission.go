package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	RoleID     int
	ResourceID int  `json:"resourceId"`
	Read       bool `json:"read"`
	Create     bool `json:"create"`
	Delete     bool `json:"delete"`
	Update     bool `json:"update"`
	View       bool `json:"view"`
}

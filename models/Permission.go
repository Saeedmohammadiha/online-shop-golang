package models

type Permission struct {
    ID    uint   `gorm:"primaryKey"`
    Title string
    Roles []Role        `gorm:"many2many:permission_roles;"`
}
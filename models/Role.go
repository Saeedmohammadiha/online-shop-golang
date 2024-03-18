package models

type Role struct {
    ID           uint           `gorm:"primaryKey"`
    Title        string
    PermissionID uint
    Users        []User        `gorm:"many2many:user_roles;"`
}
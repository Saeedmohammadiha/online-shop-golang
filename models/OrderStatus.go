package models

type OrderStatus struct {
    ID    uint   `gorm:"primaryKey"`
    Title string
}
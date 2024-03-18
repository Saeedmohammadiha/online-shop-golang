package models



type TransactionStatus struct {
    ID    uint   `gorm:"primaryKey"`
    Title string
}
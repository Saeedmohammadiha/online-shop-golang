package models

import (
	"time"
)

type Transaction struct {
    ID                  uint           `gorm:"primaryKey"`
    OrderID             uint
    UserID              uint
    Amount              float64
    Receipt             []byte         `gorm:"type:BINARY"`
    TrackingCode        int
    TransactionStatusID uint
    CreatedAt           time.Time
}
package models

import (
	"time"
)

type Transaction struct {
	ID                  uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserID              uint      `json:"userId"`
	OrderID             uint      `json:"orderId"`
	Amount              float64   `json:"amount"`
	Reciept             []byte    `json:"reciept,omitempty"`
	TrackingCode        uint      `json:"trackingCode"`
	TransactionStatusID uint      `json:"transactionStatus,omitempty"` // Optional relationship with TransactionStatus model
	CreatedAt           time.Time `json:"createdAt"`
}

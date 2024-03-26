package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID              uint    `json:"userId"`
	OrderID             uint    `json:"orderId"`
	Amount              float64 `json:"amount"`
	Reciept             []byte  `json:"reciept,omitempty"`
	TrackingCode        uint    `json:"trackingCode"`
	TransactionStatusID uint    `json:"transactionStatus,omitempty"`
}

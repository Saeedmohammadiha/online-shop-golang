package models

import (
	"time"
)

type Order struct {
	ID                 uint        `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserID             uint        `json:"userId"`
	Tax                float64     `json:"tax"`
	PostalTrackingCode uint        `json:"postalTrackingCode"`
	CreatedAt          time.Time   `json:"createdAt"`
	OrderStatusID      uint        `json:"statusId"`
	AddressID          uint        `json:"addressId"`
	Address            Address     `json:"address"`
	OrderItems         []OrderItem `json:"orderItems,omitempty"`
	Transaction        Transaction ` json:"transaction,omitempty"`
}

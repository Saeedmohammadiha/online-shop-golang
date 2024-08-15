package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID             uint        `json:"userId"`
	Tax                float64     `json:"tax"`
	PostalTrackingCode uint        `json:"postalTrackingCode"`
	OrderStatusID      uint        `json:"statusId"`
	AddressID          uint        `json:"addressId"`
	Address            Address     `json:"address"`
	OrderItems         []OrderItem `json:"orderItems,omitempty"`
	Transaction        Transaction `json:"transaction,omitempty"`
}

package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID          uint
	Name        string
	LastName    string
	Email       *string
	PhoneNumber sql.NullString
	Orders       []Order
	Address     Address
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
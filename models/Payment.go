package models

import (
	"time"
)

type Payment struct {
	ID          uint
	CreatedAt   time.Time
	UserId      User
	OrderRecipt      string
	TotalAmount float64
}

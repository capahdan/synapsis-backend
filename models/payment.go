package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	UserID      uint
	OrderID     uint
	PaymentType string
	Amount      int
}

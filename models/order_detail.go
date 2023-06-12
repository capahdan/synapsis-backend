package models

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	ProductID uint
	OrderID   uint
	Quantity  int
	SubTotal  int
	Discount  int
}

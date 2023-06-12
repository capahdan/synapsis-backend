package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID      uint
	TotalPrice  int
	Status      bool
	Payment     Payment       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderDetail []OrderDetail `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

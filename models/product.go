package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	CategoryID  uint
	Name        string
	Description string
	Price       int
	Stock       int
	Status      bool
	Carts       []Cart `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

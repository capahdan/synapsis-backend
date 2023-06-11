package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Category string
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName    string
	Email       string `gorm:"unique"`
	Password    string
	PhoneNumber string
	Gender      string
	BirthDate   *time.Time
	Citizen     string
	Carts       []Cart    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Orders      []Order   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Payments    []Payment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

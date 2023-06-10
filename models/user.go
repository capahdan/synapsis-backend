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
}
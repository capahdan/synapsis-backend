package dtos

import "time"

type ProductInput struct {
	CategoryID  uint   `json:"category_id" example:"1"`
	Name        string `json:"name" form:"name" example:"Erigo"`
	Description string `json:"description" example:"Pakaian Erigo Keluaran Terbaru"`
	Price       int    `json:"price" example:"100000"`
	Stock       int    `json:"stock" example:"100"`
	Status      bool   `json:"status" example:"true"`
}

type ProductResponse struct {
	ProductID   uint      `json:"product_id" example:"1"`
	CategoryID  uint      `json:"category_id" example:"1"`
	Name        string    `json:"name" example:"Erigo"`
	Description string    `json:"description" example:"Pakaian Erigo Keluaran Terbaru"`
	Price       int       `json:"price" example:"100000"`
	Stock       int       `json:"stock" example:"100"`
	Status      bool      `json:"status" example:"true"`
	CreatedAt   time.Time `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt   time.Time `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}

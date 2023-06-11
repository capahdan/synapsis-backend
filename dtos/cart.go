package dtos

import "time"

type CartInput struct {
	UserID    uint `json:"user_id" example:"1"`
	ProductID uint `json:"product_id" example:"1"`
	Price     int  `json:"price" example:"100000"`
	Quantity  int  `json:"quantity" example:"2"`
}

type CartResponse struct {
	CartID    uint      `json:"cart_id" example:"1"`
	UserID    uint      `json:"user_id" example:"1"`
	ProductID uint      `json:"product_id" example:"1"`
	Price     int       `json:"price" example:"100000"`
	Quantity  int       `json:"quantity" example:"2"`
	CreatedAt time.Time `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}

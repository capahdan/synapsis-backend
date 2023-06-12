package dtos

import "time"

type OrderInput struct {
	UserID     uint `json:"user_id" example:"1"`
	TotalPrice int  `json:"total_price" example:"100000"`
	Status     bool `json:"status" example:"true"`
}

type OrderResponse struct {
	OrderID    uint      `json:"order_id" example:"1"`
	UserID     uint      `json:"user_id" example:"1"`
	TotalPrice int       `json:"total_price" example:"100000"`
	Status     bool      `json:"status" example:"true"`
	CreatedAt  time.Time `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt  time.Time `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}

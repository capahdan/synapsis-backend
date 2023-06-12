package dtos

import "time"

type OrderInput struct {
	UserID     uint `json:"user_id" example:"1"`
	TotalPrice int  `json:"total_price" example:"100000"`
}

type OrderInputCheckout struct {
	UserID uint `json:"user_id" example:"1"`
}

type OrderResponse struct {
	OrderID    uint      `json:"order_id" example:"1"`
	UserID     uint      `json:"user_id" example:"1"`
	TotalPrice int       `json:"total_price" example:"100000"`
	Status     string    `json:"status" example:"unpaid"`
	CreatedAt  time.Time `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt  time.Time `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}

type OrderResponseCheckout struct {
	OrderID     uint                  `json:"order_id" example:"1"`
	UserID      uint                  `json:"user_id" example:"1"`
	TotalPrice  int                   `json:"total_price" example:"100000"`
	Status      string                `json:"status" example:"unpaid"`
	OrderDetail []OrderDetailResponse `json:"order_detail"`
	CreatedAt   time.Time             `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt   time.Time             `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}

package dtos

import "time"

type PaymentInput struct {
	OrderID     uint   `json:"order_id" example:"1"`
	UserID      uint   `json:"user_id" example:"1"`
	PaymentType string `json:"payment_type" example:"100000"`
	Amount      int    `json:"amount" example:"100000"`
}

type PaymentResponse struct {
	PaymentID   uint      `json:"payment_id" example:"1"`
	OrderID     uint      `json:"order_id" example:"1"`
	UserID      uint      `json:"user_id" example:"1"`
	PaymentType string    `json:"payment_type" example:"transfer"`
	Amount      int       `json:"amount" example:"100000"`
	CreatedAt   time.Time `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt   time.Time `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}

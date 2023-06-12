package dtos

import "time"

type OrderDetailInput struct {
	ProductID uint `json:"product_id" example:"1"`
	OrderID   uint `json:"order_id" example:"1"`
	Quantity  int  `json:"quantity" example:"2"`
	SubTotal  int  `json:"sub_total" example:"200000"`
	Discount  int  `json:"discount" example:"0"`
}

type OrderDetailResponse struct {
	OrderDetailID uint      `json:"order_detail_id" example:"1"`
	ProductID     uint      `json:"product_id" example:"1"`
	OrderID       uint      `json:"order_id" example:"1"`
	Quantity      int       `json:"quantity" example:"2"`
	SubTotal      int       `json:"sub_total" example:"200000"`
	Discount      int       `json:"discount" example:"0"`
	CreatedAt     time.Time `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt     time.Time `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}

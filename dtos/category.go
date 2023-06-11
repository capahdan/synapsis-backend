package dtos

import "time"

type CategoryInput struct {
	Category string `json:"category" form:"category" example:"pakaian"`
}

type CategoryResponse struct {
	CategoryID uint      `json:"category_id" example:"1"`
	Category   string    `json:"category" example:"pakaian"`
	CreatedAt  time.Time `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt  time.Time `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}

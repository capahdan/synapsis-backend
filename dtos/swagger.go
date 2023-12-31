package dtos

import "synapsis-backend/helpers"

type UserStatusOKResponse struct {
	StatusCode int                     `json:"status_code" example:"200"`
	Message    string                  `json:"message" example:"Successfully get user credentials"`
	Data       UserInformationResponse `json:"data"`
}

type UserCreatedResponse struct {
	StatusCode int                     `json:"status_code" example:"201"`
	Message    string                  `json:"message" example:"Successfully registered"`
	Data       UserInformationResponse `json:"data"`
}
type CategoryCreatedResponse struct {
	StatusCode int              `json:"status_code" example:"201"`
	Message    string           `json:"message" example:"Successfully created category"`
	Data       CategoryResponse `json:"data"`
}

type GetAllCategoryStatusOKResponse struct {
	StatusCode int              `json:"status_code" example:"200"`
	Message    string           `json:"message" example:"Successfully get category"`
	Data       CategoryResponse `json:"data"`
	Meta       helpers.Meta     `json:"meta"`
}

type CategoryStatusOKResponse struct {
	StatusCode int              `json:"status_code" example:"200"`
	Message    string           `json:"message" example:"Successfully get category"`
	Data       CategoryResponse `json:"data"`
}
type ProductCreatedResponse struct {
	StatusCode int             `json:"status_code" example:"201"`
	Message    string          `json:"message" example:"Successfully created product"`
	Data       ProductResponse `json:"data"`
}

type GetAllProductStatusOKResponse struct {
	StatusCode int             `json:"status_code" example:"200"`
	Message    string          `json:"message" example:"Successfully get product"`
	Data       ProductResponse `json:"data"`
	Meta       helpers.Meta    `json:"meta"`
}
type ProductStatusOKResponse struct {
	StatusCode int             `json:"status_code" example:"200"`
	Message    string          `json:"message" example:"Successfully get product"`
	Data       ProductResponse `json:"data"`
}
type CartCreatedResponse struct {
	StatusCode int          `json:"status_code" example:"201"`
	Message    string       `json:"message" example:"Successfully created cart"`
	Data       CartResponse `json:"data"`
}

type GetAllCartStatusOKResponse struct {
	StatusCode int          `json:"status_code" example:"200"`
	Message    string       `json:"message" example:"Successfully get cart"`
	Data       CartResponse `json:"data"`
	Meta       helpers.Meta `json:"meta"`
}
type CartStatusOKResponse struct {
	StatusCode int          `json:"status_code" example:"200"`
	Message    string       `json:"message" example:"Successfully get cart"`
	Data       CartResponse `json:"data"`
}
type OrderCreatedResponse struct {
	StatusCode int           `json:"status_code" example:"201"`
	Message    string        `json:"message" example:"Successfully created order"`
	Data       OrderResponse `json:"data"`
}

type GetAllOrderStatusOKResponse struct {
	StatusCode int           `json:"status_code" example:"200"`
	Message    string        `json:"message" example:"Successfully get order"`
	Data       OrderResponse `json:"data"`
	Meta       helpers.Meta  `json:"meta"`
}
type OrderStatusOKResponse struct {
	StatusCode int           `json:"status_code" example:"200"`
	Message    string        `json:"message" example:"Successfully get order"`
	Data       OrderResponse `json:"data"`
}
type OrderCheckoutStatusOKResponse struct {
	StatusCode int                     `json:"status_code" example:"200"`
	Message    string                  `json:"message" example:"Successfully get order"`
	Data       []OrderResponseCheckout `json:"data"`
}
type PaymentCreatedResponse struct {
	StatusCode int             `json:"status_code" example:"201"`
	Message    string          `json:"message" example:"Successfully created payment"`
	Data       PaymentResponse `json:"data"`
}

type GetAllPaymentStatusOKResponse struct {
	StatusCode int             `json:"status_code" example:"200"`
	Message    string          `json:"message" example:"Successfully get payment"`
	Data       PaymentResponse `json:"data"`
	Meta       helpers.Meta    `json:"meta"`
}
type PaymentStatusOKResponse struct {
	StatusCode int             `json:"status_code" example:"200"`
	Message    string          `json:"message" example:"Successfully get payment"`
	Data       PaymentResponse `json:"data"`
}
type OrderDetailCreatedResponse struct {
	StatusCode int                 `json:"status_code" example:"201"`
	Message    string              `json:"message" example:"Successfully created order"`
	Data       OrderDetailResponse `json:"data"`
}

type GetAllOrderDetailStatusOKResponse struct {
	StatusCode int                 `json:"status_code" example:"200"`
	Message    string              `json:"message" example:"Successfully get order"`
	Data       OrderDetailResponse `json:"data"`
	Meta       helpers.Meta        `json:"meta"`
}
type OrderDetailStatusOKResponse struct {
	StatusCode int                 `json:"status_code" example:"200"`
	Message    string              `json:"message" example:"Successfully get order"`
	Data       OrderDetailResponse `json:"data"`
}

type StatusOKDeletedResponse struct {
	StatusCode int         `json:"status_code" example:"200"`
	Message    string      `json:"message" example:"Successfully deleted"`
	Errors     interface{} `json:"errors"`
}

type BadRequestResponse struct {
	StatusCode int         `json:"status_code" example:"400"`
	Message    string      `json:"message" example:"Bad Request"`
	Errors     interface{} `json:"errors"`
}

type UnauthorizedResponse struct {
	StatusCode int         `json:"status_code" example:"401"`
	Message    string      `json:"message" example:"Unauthorized"`
	Errors     interface{} `json:"errors"`
}

type ForbiddenResponse struct {
	StatusCode int         `json:"status_code" example:"403"`
	Message    string      `json:"message" example:"Forbidden"`
	Errors     interface{} `json:"errors"`
}

type NotFoundResponse struct {
	StatusCode int         `json:"status_code" example:"404"`
	Message    string      `json:"message" example:"Not Found"`
	Errors     interface{} `json:"errors"`
}

type InternalServerErrorResponse struct {
	StatusCode int         `json:"status_code" example:"500"`
	Message    string      `json:"message" example:"Internal Server Error"`
	Errors     interface{} `json:"errors"`
}

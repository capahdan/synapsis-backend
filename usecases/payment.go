package usecases

import (
	"synapsis-backend/dtos"
	"synapsis-backend/models"
	"synapsis-backend/repositories"
)

type PaymentUsecase interface {
	GetAllPayments(page, limit, user_id int) ([]dtos.PaymentResponse, int, error)
	GetPaymentByID(id uint) (dtos.PaymentResponse, error)
	CreatePayment(payment *dtos.PaymentInput) (dtos.PaymentResponse, error)
	UpdatePayment(id uint, paymentInput dtos.PaymentInput) (dtos.PaymentResponse, error)
	DeletePayment(id uint) error
}

type paymentUsecase struct {
	paymentRepo repositories.PaymentRepository
}

func NewPaymentUsecase(PaymentRepo repositories.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{PaymentRepo}
}

// GetAllPayments godoc
// @Summary      Get all payment
// @Description  Get all payment
// @Tags         Payment
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param user_id query int false "Seacrh by category ID"
// @Success      200 {object} dtos.GetAllPaymentStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /payment [get]
// @Security BearerAuth
func (u *paymentUsecase) GetAllPayments(page, limit, user_id int) ([]dtos.PaymentResponse, int, error) {
	payments, count, err := u.paymentRepo.GetAllPayments(page, limit, user_id)
	if err != nil {
		return nil, 0, err
	}

	var paymentResponses []dtos.PaymentResponse
	for _, payment := range payments {
		// category, err := u.paymentRepo.GetCategoryByID(payment.CategoryID)

		paymentResponse := dtos.PaymentResponse{
			PaymentID:   payment.ID,
			OrderID:     payment.OrderID,
			UserID:      payment.UserID,
			PaymentType: payment.PaymentType,
			Amount:      payment.Amount,
			CreatedAt:   payment.CreatedAt,
			UpdatedAt:   payment.UpdatedAt,
		}
		paymentResponses = append(paymentResponses, paymentResponse)
	}

	return paymentResponses, count, nil
}

// GetPaymentByID godoc
// @Summary      Get payment by ID
// @Description  Get payment by ID
// @Tags         Payment
// @Accept       json
// @Produce      json
// @Param id path integer true "ID payment"
// @Success      200 {object} dtos.PaymentStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /payment/{id} [get]
// @Security BearerAuth
func (u *paymentUsecase) GetPaymentByID(id uint) (dtos.PaymentResponse, error) {
	var paymentResponses dtos.PaymentResponse
	payment, err := u.paymentRepo.GetPaymentByID(id)
	if err != nil {
		return paymentResponses, err
	}
	paymentResponse := dtos.PaymentResponse{
		PaymentID:   payment.ID,
		OrderID:     payment.OrderID,
		UserID:      payment.UserID,
		PaymentType: payment.PaymentType,
		Amount:      payment.Amount,
		CreatedAt:   payment.CreatedAt,
		UpdatedAt:   payment.UpdatedAt,
	}
	return paymentResponse, nil
}

// CreatePayment godoc
// @Summary      Create a new payment
// @Description  Create a new payment
// @Tags         Payment
// @Accept       json
// @Produce      json
// @Param        request body dtos.PaymentInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.PaymentStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /payment [post]
// @Security BearerAuth
func (u *paymentUsecase) CreatePayment(payment *dtos.PaymentInput) (dtos.PaymentResponse, error) {
	var paymentResponses dtos.PaymentResponse

	createPayment := models.Payment{
		UserID:      payment.UserID,
		OrderID:     payment.OrderID,
		PaymentType: payment.PaymentType,
		Amount:      payment.Amount,
	}

	createdPayment, err := u.paymentRepo.CreatePayment(createPayment)
	if err != nil {
		return paymentResponses, err
	}

	paymentResponse := dtos.PaymentResponse{
		PaymentID:   createdPayment.ID,
		OrderID:     createdPayment.OrderID,
		UserID:      createdPayment.UserID,
		PaymentType: createdPayment.PaymentType,
		Amount:      createdPayment.Amount,
		CreatedAt:   createdPayment.CreatedAt,
		UpdatedAt:   createdPayment.UpdatedAt,
	}

	return paymentResponse, nil
}

// UpdatePayment godoc
// @Summary      Update payment
// @Description  Update payment
// @Tags         Payment
// @Accept       json
// @Produce      json
// @Param id path integer true "ID payment"
// @Param        request body dtos.PaymentInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.PaymentStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /payment/{id} [put]
// @Security BearerAuth
func (u *paymentUsecase) UpdatePayment(id uint, paymentInput dtos.PaymentInput) (dtos.PaymentResponse, error) {

	var payment models.Payment
	var paymentResponse dtos.PaymentResponse

	payment, err := u.paymentRepo.GetPaymentByID(id)
	if err != nil {
		return paymentResponse, err
	}

	payment.ID = id
	payment.UserID = paymentInput.UserID
	payment.OrderID = paymentInput.OrderID
	payment.PaymentType = paymentInput.PaymentType
	payment.Amount = paymentInput.Amount

	payment, err = u.paymentRepo.UpdatePayment(payment)

	if err != nil {
		return paymentResponse, err
	}

	paymentResponse.PaymentID = payment.ID
	paymentResponse.UserID = payment.UserID
	paymentResponse.OrderID = payment.OrderID
	paymentResponse.PaymentType = payment.PaymentType
	paymentResponse.Amount = payment.Amount
	paymentResponse.CreatedAt = payment.CreatedAt
	paymentResponse.UpdatedAt = payment.UpdatedAt

	return paymentResponse, nil

}

// DeletePayment godoc
// @Summary      Delete a payment
// @Description  Delete a payment
// @Tags         Payment
// @Accept       json
// @Produce      json
// @Param id path integer true "ID payment"
// @Success      200 {object} dtos.StatusOKDeletedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /payment/{id} [delete]
// @Security BearerAuth
func (u *paymentUsecase) DeletePayment(id uint) error {
	payment, err := u.paymentRepo.GetPaymentByID(id)

	if err != nil {
		return nil
	}
	err = u.paymentRepo.DeletePayment(payment)
	return err
}

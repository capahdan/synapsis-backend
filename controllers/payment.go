package controllers

import (
	"net/http"
	"strconv"
	"synapsis-backend/dtos"
	"synapsis-backend/helpers"
	"synapsis-backend/usecases"

	"github.com/labstack/echo/v4"
)

type PaymentController interface {
	GetAllPayments(c echo.Context) error
	GetPaymentByID(c echo.Context) error
	CreatePayment(c echo.Context) error
	UpdatePayment(c echo.Context) error
	DeletePayment(c echo.Context) error
}

type paymentController struct {
	paymentUsecase usecases.PaymentUsecase
}

func NewPaymentController(paymentUsecase usecases.PaymentUsecase) PaymentController {
	return &paymentController{paymentUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *paymentController) GetAllPayments(ctx echo.Context) error {
	pageParam := ctx.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		page = 1
	}

	limitParam := ctx.QueryParam("limit")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		limit = 10
	}
	userParam := ctx.QueryParam("user_id")
	user_id, err := strconv.Atoi(userParam)

	payments, count, err := c.paymentUsecase.GetAllPayments(page, limit, user_id)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get all payment",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully get all payments",
			payments,
			page,
			limit,
			count,
		),
	)
}

func (c *paymentController) GetPaymentByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	payment, err := c.paymentUsecase.GetPaymentByID(uint(id))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get payment by id",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get payment by id",
			payment,
		),
	)

}

func (c *paymentController) CreatePayment(ctx echo.Context) error {
	var paymentDTO dtos.PaymentInput
	if err := ctx.Bind(&paymentDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	payment, err := c.paymentUsecase.CreatePayment(&paymentDTO)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a payment",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a payment",
			payment,
		),
	)
}

func (c *paymentController) UpdatePayment(ctx echo.Context) error {

	var paymentInput dtos.PaymentInput
	if err := ctx.Bind(&paymentInput); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	payment, err := c.paymentUsecase.GetPaymentByID(uint(id))
	if payment.PaymentID == 0 {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get payment by id",
				helpers.GetErrorData(err),
			),
		)
	}

	paymentResp, err := c.paymentUsecase.UpdatePayment(uint(id), paymentInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to updated a payment",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated payment",
			paymentResp,
		),
	)
}

func (c *paymentController) DeletePayment(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.paymentUsecase.DeletePayment(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully deleted payment",
			nil,
		),
	)
}

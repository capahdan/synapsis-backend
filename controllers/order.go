package controllers

import (
	"net/http"
	"strconv"
	"synapsis-backend/dtos"
	"synapsis-backend/helpers"
	"synapsis-backend/usecases"

	"github.com/labstack/echo/v4"
)

type OrderController interface {
	GetAllOrders(c echo.Context) error
	GetOrderByID(c echo.Context) error
	CreateOrder(c echo.Context) error
	UpdateOrder(c echo.Context) error
	DeleteOrder(c echo.Context) error
	Checkout(c echo.Context) error
}

type orderController struct {
	orderUsecase usecases.OrderUsecase
}

func NewOrderController(orderUsecase usecases.OrderUsecase) OrderController {
	return &orderController{orderUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *orderController) GetAllOrders(ctx echo.Context) error {
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

	orders, count, err := c.orderUsecase.GetAllOrders(page, limit, user_id)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get all order",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully get all orders",
			orders,
			page,
			limit,
			count,
		),
	)
}

func (c *orderController) GetOrderByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	order, err := c.orderUsecase.GetOrderByID(uint(id))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get order by id",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get order by id",
			order,
		),
	)

}

func (c *orderController) CreateOrder(ctx echo.Context) error {
	var orderDTO dtos.OrderInput
	if err := ctx.Bind(&orderDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	order, err := c.orderUsecase.CreateOrder(&orderDTO)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a order",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a order",
			order,
		),
	)
}

func (c *orderController) Checkout(ctx echo.Context) error {
	var orderDTO dtos.OrderInputCheckout
	if err := ctx.Bind(&orderDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	order, err := c.orderUsecase.Checkout(&orderDTO)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a order",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a order",
			order,
		),
	)
}

func (c *orderController) UpdateOrder(ctx echo.Context) error {

	var orderInput dtos.OrderInput
	if err := ctx.Bind(&orderInput); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	order, err := c.orderUsecase.GetOrderByID(uint(id))
	if order.OrderID == 0 {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get order by id",
				helpers.GetErrorData(err),
			),
		)
	}

	orderResp, err := c.orderUsecase.UpdateOrder(uint(id), orderInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to updated a order",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated order",
			orderResp,
		),
	)
}

func (c *orderController) DeleteOrder(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.orderUsecase.DeleteOrder(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully deleted order",
			nil,
		),
	)
}

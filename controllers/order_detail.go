package controllers

import (
	"net/http"
	"strconv"
	"synapsis-backend/dtos"
	"synapsis-backend/helpers"
	"synapsis-backend/usecases"

	"github.com/labstack/echo/v4"
)

type OrderDetailController interface {
	GetAllOrderDetails(c echo.Context) error
	GetOrderDetailByID(c echo.Context) error
	CreateOrderDetail(c echo.Context) error
	UpdateOrderDetail(c echo.Context) error
	DeleteOrderDetail(c echo.Context) error
}

type orderDetailController struct {
	orderDetailUsecase usecases.OrderDetailUsecase
}

func NewOrderDetailController(orderDetailUsecase usecases.OrderDetailUsecase) OrderDetailController {
	return &orderDetailController{orderDetailUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *orderDetailController) GetAllOrderDetails(ctx echo.Context) error {
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

	orderDetails, count, err := c.orderDetailUsecase.GetAllOrderDetails(page, limit, user_id)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get all orderDetail",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully get all orderDetails",
			orderDetails,
			page,
			limit,
			count,
		),
	)
}

func (c *orderDetailController) GetOrderDetailByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	orderDetail, err := c.orderDetailUsecase.GetOrderDetailByID(uint(id))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get orderDetail by id",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get orderDetail by id",
			orderDetail,
		),
	)

}

func (c *orderDetailController) CreateOrderDetail(ctx echo.Context) error {
	var orderDetailDTO dtos.OrderDetailInput
	if err := ctx.Bind(&orderDetailDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	orderDetail, err := c.orderDetailUsecase.CreateOrderDetail(&orderDetailDTO)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a orderDetail",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a orderDetail",
			orderDetail,
		),
	)
}

func (c *orderDetailController) UpdateOrderDetail(ctx echo.Context) error {

	var orderDetailInput dtos.OrderDetailInput
	if err := ctx.Bind(&orderDetailInput); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	orderDetail, err := c.orderDetailUsecase.GetOrderDetailByID(uint(id))
	if orderDetail.OrderDetailID == 0 {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get orderDetail by id",
				helpers.GetErrorData(err),
			),
		)
	}

	orderDetailResp, err := c.orderDetailUsecase.UpdateOrderDetail(uint(id), orderDetailInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to updated a orderDetail",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated orderDetail",
			orderDetailResp,
		),
	)
}

func (c *orderDetailController) DeleteOrderDetail(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.orderDetailUsecase.DeleteOrderDetail(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully deleted orderDetail",
			nil,
		),
	)
}

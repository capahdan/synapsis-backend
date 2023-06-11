package controllers

import (
	"net/http"
	"strconv"
	"synapsis-backend/dtos"
	"synapsis-backend/helpers"
	"synapsis-backend/usecases"

	"github.com/labstack/echo/v4"
)

type CartController interface {
	GetAllCarts(c echo.Context) error
	GetCartByID(c echo.Context) error
	CreateCart(c echo.Context) error
	UpdateCart(c echo.Context) error
	DeleteCart(c echo.Context) error
}

type cartController struct {
	cartUsecase usecases.CartUsecase
}

func NewCartController(cartUsecase usecases.CartUsecase) CartController {
	return &cartController{cartUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *cartController) GetAllCarts(ctx echo.Context) error {
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

	user_id_param := ctx.QueryParam("user_id")
	user_id, err := strconv.Atoi(user_id_param)
	if err != nil {
		user_id = 0
	}

	carts, count, err := c.cartUsecase.GetAllCarts(page, limit, user_id)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get all cart",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully get all carts",
			carts,
			page,
			limit,
			count,
		),
	)
}

func (c *cartController) GetCartByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	cart, err := c.cartUsecase.GetCartByID(uint(id))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get cart by id",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get cart by id",
			cart,
		),
	)

}

func (c *cartController) CreateCart(ctx echo.Context) error {
	var cartDTO dtos.CartInput
	if err := ctx.Bind(&cartDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	cart, err := c.cartUsecase.CreateCart(&cartDTO)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a cart",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a cart",
			cart,
		),
	)
}

func (c *cartController) UpdateCart(ctx echo.Context) error {

	var cartInput dtos.CartInput
	if err := ctx.Bind(&cartInput); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	cart, err := c.cartUsecase.GetCartByID(uint(id))
	if cart.CartID == 0 {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get cart by id",
				helpers.GetErrorData(err),
			),
		)
	}

	cartResp, err := c.cartUsecase.UpdateCart(uint(id), cartInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to updated a cart",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated cart",
			cartResp,
		),
	)
}

func (c *cartController) DeleteCart(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.cartUsecase.DeleteCart(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully deleted cart",
			nil,
		),
	)
}

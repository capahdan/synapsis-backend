package controllers

import (
	"net/http"
	"strconv"
	"synapsis-backend/dtos"
	"synapsis-backend/helpers"
	"synapsis-backend/usecases"

	"github.com/labstack/echo/v4"
)

type ProductController interface {
	GetAllProducts(c echo.Context) error
	GetProductByID(c echo.Context) error
	CreateProduct(c echo.Context) error
	UpdateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
}

type productController struct {
	productUsecase usecases.ProductUsecase
}

func NewProductController(productUsecase usecases.ProductUsecase) ProductController {
	return &productController{productUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *productController) GetAllProducts(ctx echo.Context) error {
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

	category_id_param := ctx.QueryParam("category_id")
	category_id, err := strconv.Atoi(category_id_param)
	if err != nil {
		category_id = 0
	}

	products, count, err := c.productUsecase.GetAllProducts(page, limit, category_id)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get all product",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully get all products",
			products,
			page,
			limit,
			count,
		),
	)
}

func (c *productController) GetProductByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	product, err := c.productUsecase.GetProductByID(uint(id))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get product by id",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get product by id",
			product,
		),
	)

}

func (c *productController) CreateProduct(ctx echo.Context) error {
	var productDTO dtos.ProductInput
	if err := ctx.Bind(&productDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	product, err := c.productUsecase.CreateProduct(&productDTO)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a product",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a product",
			product,
		),
	)
}

func (c *productController) UpdateProduct(ctx echo.Context) error {

	var productInput dtos.ProductInput
	if err := ctx.Bind(&productInput); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	product, err := c.productUsecase.GetProductByID(uint(id))
	if product.ProductID == 0 {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get product by id",
				helpers.GetErrorData(err),
			),
		)
	}

	productResp, err := c.productUsecase.UpdateProduct(uint(id), productInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to updated a product",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated product",
			productResp,
		),
	)
}

func (c *productController) DeleteProduct(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.productUsecase.DeleteProduct(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully deleted product",
			nil,
		),
	)
}

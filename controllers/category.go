package controllers

import (
	"net/http"
	"strconv"
	"synapsis-backend/dtos"
	"synapsis-backend/helpers"
	"synapsis-backend/usecases"

	"github.com/labstack/echo/v4"
)

type CategoryController interface {
	GetAllCategorys(c echo.Context) error
	GetCategoryByID(c echo.Context) error
	CreateCategory(c echo.Context) error
	UpdateCategory(c echo.Context) error
	DeleteCategory(c echo.Context) error
}

type categoryController struct {
	categoryUsecase usecases.CategoryUsecase
}

func NewCategoryController(categoryUsecase usecases.CategoryUsecase) CategoryController {
	return &categoryController{categoryUsecase}
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *categoryController) GetAllCategorys(ctx echo.Context) error {
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

	categorys, count, err := c.categoryUsecase.GetAllCategorys(page, limit)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get all category",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewPaginationResponse(
			http.StatusOK,
			"Successfully get all categorys",
			categorys,
			page,
			limit,
			count,
		),
	)
}

func (c *categoryController) GetCategoryByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	category, err := c.categoryUsecase.GetCategoryByID(uint(id))

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get category by id",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully to get category by id",
			category,
		),
	)

}

func (c *categoryController) CreateCategory(ctx echo.Context) error {
	var categoryDTO dtos.CategoryInput
	if err := ctx.Bind(&categoryDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	category, err := c.categoryUsecase.CreateCategory(&categoryDTO)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to created a category",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully to created a category",
			category,
		),
	)
}

func (c *categoryController) UpdateCategory(ctx echo.Context) error {

	var categoryInput dtos.CategoryInput
	if err := ctx.Bind(&categoryInput); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	category, err := c.categoryUsecase.GetCategoryByID(uint(id))
	if category.CategoryID == 0 {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get category by id",
				helpers.GetErrorData(err),
			),
		)
	}

	categoryResp, err := c.categoryUsecase.UpdateCategory(uint(id), categoryInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to updated a category",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated category",
			categoryResp,
		),
	)
}

func (c *categoryController) DeleteCategory(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.categoryUsecase.DeleteCategory(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully deleted category",
			nil,
		),
	)
}

package usecases

import (
	"synapsis-backend/dtos"
	"synapsis-backend/repositories"
)

type CategoryUsecase interface {
	GetAllCategorys(page, limit int) ([]dtos.CategoryResponse, int, error)
	GetCategoryByID(id uint) (dtos.CategoryResponse, error)
	CreateCategory(category *dtos.CategoryInput) (dtos.CategoryResponse, error)
	UpdateCategory(id uint, categoryInput dtos.CategoryInput) (dtos.CategoryResponse, error)
	DeleteCategory(id uint) error
}

type categoryUsecase struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryUsecase(CategoryRepo repositories.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{CategoryRepo}
}

// GetAllCategorys godoc
// @Summary      Get all category
// @Description  Get all category
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success      200 {object} dtos.GetAllCategoryStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/category [get]
// @Security BearerAuth
func (u *categoryUsecase) GetAllCategorys(page, limit int) ([]dtos.CategoryResponse, int, error) {
	return nil, 0, nil
}

// GetCategoryByID godoc
// @Summary      Get category by ID
// @Description  Get category by ID
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param id path integer true "ID category"
// @Success      200 {object} dtos.CategoryStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/category/{id} [get]
// @Security BearerAuth
func (u *categoryUsecase) GetCategoryByID(id uint) (dtos.CategoryResponse, error) {
	return dtos.CategoryResponse{}, nil
}

// CreateCategory godoc
// @Summary      Create a new category
// @Description  Create a new category
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param        request body dtos.CategoryInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.CategoryStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/category [post]
// @Security BearerAuth
func (u *categoryUsecase) CreateCategory(category *dtos.CategoryInput) (dtos.CategoryResponse, error) {

	return dtos.CategoryResponse{}, nil
}

// UpdateCategory godoc
// @Summary      Update category
// @Description  Update category
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param id path integer true "ID category"
// @Param        request body dtos.CategoryInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.CategoryStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/category [put]
// @Security BearerAuth
func (u *categoryUsecase) UpdateCategory(id uint, categoryInput dtos.CategoryInput) (dtos.CategoryResponse, error) {

	return dtos.CategoryResponse{}, nil

}

// DeleteCategory godoc
// @Summary      Delete a category
// @Description  Delete a category
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param id path integer true "ID category"
// @Success      200 {object} dtos.StatusOKDeletedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/category/{id} [delete]
// @Security BearerAuth
func (u *categoryUsecase) DeleteCategory(id uint) error {
	category, err := u.categoryRepo.GetCategoryByID(id)

	if err != nil {
		return nil
	}
	err = u.categoryRepo.DeleteCategory(category)
	return err
}

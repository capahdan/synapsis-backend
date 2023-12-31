package usecases

import (
	"fmt"
	"synapsis-backend/dtos"
	"synapsis-backend/models"
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
// @Router       /category [get]
// @Security BearerAuth
func (u *categoryUsecase) GetAllCategorys(page, limit int) ([]dtos.CategoryResponse, int, error) {
	categorys, count, err := u.categoryRepo.GetAllCategorys(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var categoryResponses []dtos.CategoryResponse
	for _, category := range categorys {
		categoryResponse := dtos.CategoryResponse{
			CategoryID: category.ID,
			Category:   category.Category,
			CreatedAt:  category.CreatedAt,
			UpdatedAt:  category.UpdatedAt,
		}
		categoryResponses = append(categoryResponses, categoryResponse)
	}

	return categoryResponses, count, nil
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
// @Router       /category/{id} [get]
// @Security BearerAuth
func (u *categoryUsecase) GetCategoryByID(id uint) (dtos.CategoryResponse, error) {
	var categoryResponses dtos.CategoryResponse
	category, err := u.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return categoryResponses, err
	}
	categoryResponse := dtos.CategoryResponse{
		CategoryID: category.ID,
		Category:   category.Category,
		CreatedAt:  category.CreatedAt,
		UpdatedAt:  category.UpdatedAt,
	}
	return categoryResponse, nil
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
// @Router       /category [post]
// @Security BearerAuth
func (u *categoryUsecase) CreateCategory(category *dtos.CategoryInput) (dtos.CategoryResponse, error) {
	var categoryResponses dtos.CategoryResponse

	createCategory := models.Category{
		Category: category.Category,
	}

	createdTrain, err := u.categoryRepo.CreateCategory(createCategory)
	if err != nil {
		return categoryResponses, err
	}

	categoryResponse := dtos.CategoryResponse{
		CategoryID: createdTrain.ID,
		Category:   createdTrain.Category,
		CreatedAt:  createdTrain.CreatedAt,
		UpdatedAt:  createdTrain.UpdatedAt,
	}

	return categoryResponse, nil
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
// @Router       /category/{id} [put]
// @Security BearerAuth
func (u *categoryUsecase) UpdateCategory(id uint, categoryInput dtos.CategoryInput) (dtos.CategoryResponse, error) {

	var category models.Category
	var categoryResponse dtos.CategoryResponse

	category, err := u.categoryRepo.GetCategoryByID(id)
	fmt.Println(category)
	if err != nil {
		return categoryResponse, err
	}

	category.Category = categoryInput.Category

	category, err = u.categoryRepo.UpdateCategory(category)

	if err != nil {
		return categoryResponse, err
	}

	categoryResponse.CategoryID = category.ID
	categoryResponse.Category = category.Category
	categoryResponse.CreatedAt = category.CreatedAt
	categoryResponse.UpdatedAt = category.UpdatedAt

	return categoryResponse, nil

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
// @Router       /category/{id} [delete]
// @Security BearerAuth
func (u *categoryUsecase) DeleteCategory(id uint) error {
	category, err := u.categoryRepo.GetCategoryByID(id)

	if err != nil {
		return nil
	}
	err = u.categoryRepo.DeleteCategory(category)
	return err
}

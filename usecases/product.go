package usecases

import (
	"synapsis-backend/dtos"
	"synapsis-backend/models"
	"synapsis-backend/repositories"
)

type ProductUsecase interface {
	GetAllProducts(page, limit, category_id int) ([]dtos.ProductResponse, int, error)
	GetProductByID(id uint) (dtos.ProductResponse, error)
	CreateProduct(product *dtos.ProductInput) (dtos.ProductResponse, error)
	UpdateProduct(id uint, productInput dtos.ProductInput) (dtos.ProductResponse, error)
	DeleteProduct(id uint) error
}

type productUsecase struct {
	productRepo repositories.ProductRepository
}

func NewProductUsecase(ProductRepo repositories.ProductRepository) ProductUsecase {
	return &productUsecase{ProductRepo}
}

// GetAllProducts godoc
// @Summary      Get all product
// @Description  Get all product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param category_id query int false "Seacrh by category ID"
// @Success      200 {object} dtos.GetAllProductStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /product [get]
// @Security BearerAuth
func (u *productUsecase) GetAllProducts(page, limit, category_id int) ([]dtos.ProductResponse, int, error) {
	products, count, err := u.productRepo.GetAllProducts(page, limit, category_id)
	if err != nil {
		return nil, 0, err
	}

	var productResponses []dtos.ProductResponse
	for _, product := range products {
		// category, err := u.productRepo.GetCategoryByID(product.CategoryID)

		productResponse := dtos.ProductResponse{
			ProductID:   product.ID,
			CategoryID:  product.CategoryID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			Stock:       product.Stock,
			Status:      product.Status,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, count, nil
}

// GetProductByID godoc
// @Summary      Get product by ID
// @Description  Get product by ID
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param id path integer true "ID product"
// @Success      200 {object} dtos.ProductStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /product/{id} [get]
// @Security BearerAuth
func (u *productUsecase) GetProductByID(id uint) (dtos.ProductResponse, error) {
	var productResponses dtos.ProductResponse
	product, err := u.productRepo.GetProductByID(id)
	if err != nil {
		return productResponses, err
	}
	productResponse := dtos.ProductResponse{
		ProductID:   product.ID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Stock:       product.Stock,
		Status:      product.Status,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
	return productResponse, nil
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Create a new product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        request body dtos.ProductInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.ProductStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /product [post]
// @Security BearerAuth
func (u *productUsecase) CreateProduct(product *dtos.ProductInput) (dtos.ProductResponse, error) {
	var productResponses dtos.ProductResponse

	createProduct := models.Product{
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Status:      product.Status,
	}

	createdProduct, err := u.productRepo.CreateProduct(createProduct)
	if err != nil {
		return productResponses, err
	}

	productResponse := dtos.ProductResponse{
		ProductID:  createdProduct.ID,
		CategoryID: createdProduct.CategoryID,
		Name:       createdProduct.Name,
		Price:      createdProduct.Price,
		Stock:      createdProduct.Stock,
		Status:     createdProduct.Status,
		CreatedAt:  createdProduct.CreatedAt,
		UpdatedAt:  createdProduct.UpdatedAt,
	}

	return productResponse, nil
}

// UpdateProduct godoc
// @Summary      Update product
// @Description  Update product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param id path integer true "ID product"
// @Param        request body dtos.ProductInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.ProductStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /product [put]
// @Security BearerAuth
func (u *productUsecase) UpdateProduct(id uint, productInput dtos.ProductInput) (dtos.ProductResponse, error) {

	var product models.Product
	var productResponse dtos.ProductResponse

	product, err := u.productRepo.GetProductByID(id)
	if err != nil {
		return productResponse, err
	}

	product.ID = id
	product.CategoryID = productInput.CategoryID
	product.Name = productInput.Name
	product.Description = productInput.Description
	product.Price = productInput.Price
	product.Stock = productInput.Stock
	product.Status = productInput.Status

	product, err = u.productRepo.UpdateProduct(product)

	if err != nil {
		return productResponse, err
	}

	productResponse.ProductID = product.ID
	productResponse.CategoryID = product.CategoryID
	productResponse.Name = product.Name
	productResponse.Description = product.Description
	productResponse.Price = product.Price
	productResponse.Stock = product.Stock
	productResponse.Status = product.Status
	productResponse.CreatedAt = product.CreatedAt
	productResponse.UpdatedAt = product.UpdatedAt

	return productResponse, nil

}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete a product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param id path integer true "ID product"
// @Success      200 {object} dtos.StatusOKDeletedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /product/{id} [delete]
// @Security BearerAuth
func (u *productUsecase) DeleteProduct(id uint) error {
	product, err := u.productRepo.GetProductByID(id)

	if err != nil {
		return nil
	}
	err = u.productRepo.DeleteProduct(product)
	return err
}

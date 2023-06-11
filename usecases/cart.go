package usecases

import (
	"synapsis-backend/dtos"
	"synapsis-backend/models"
	"synapsis-backend/repositories"
)

type CartUsecase interface {
	GetAllCarts(page, limit, user_id int) ([]dtos.CartResponse, int, error)
	GetCartByID(id uint) (dtos.CartResponse, error)
	CreateCart(cart *dtos.CartInput) (dtos.CartResponse, error)
	UpdateCart(id uint, cartInput dtos.CartInput) (dtos.CartResponse, error)
	DeleteCart(id uint) error
}

type cartUsecase struct {
	cartRepo repositories.CartRepository
}

func NewCartUsecase(CartRepo repositories.CartRepository) CartUsecase {
	return &cartUsecase{CartRepo}
}

// GetAllCarts godoc
// @Summary      Get all cart
// @Description  Get all cart
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param category_id query int false "Seacrh by category ID"
// @Success      200 {object} dtos.GetAllCartStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /cart [get]
// @Security BearerAuth
func (u *cartUsecase) GetAllCarts(page, limit, user_id int) ([]dtos.CartResponse, int, error) {
	carts, count, err := u.cartRepo.GetAllCarts(page, limit, user_id)
	if err != nil {
		return nil, 0, err
	}

	var cartResponses []dtos.CartResponse
	for _, cart := range carts {
		// category, err := u.cartRepo.GetCategoryByID(cart.CategoryID)

		cartResponse := dtos.CartResponse{
			CartID:    cart.ID,
			UserID:    cart.UserID,
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			Price:     cart.Price,
			CreatedAt: cart.CreatedAt,
			UpdatedAt: cart.UpdatedAt,
		}
		cartResponses = append(cartResponses, cartResponse)
	}

	return cartResponses, count, nil
}

// GetCartByID godoc
// @Summary      Get cart by ID
// @Description  Get cart by ID
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Param id path integer true "ID cart"
// @Success      200 {object} dtos.CartStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /cart/{id} [get]
// @Security BearerAuth
func (u *cartUsecase) GetCartByID(id uint) (dtos.CartResponse, error) {
	var cartResponses dtos.CartResponse
	cart, err := u.cartRepo.GetCartByID(id)
	if err != nil {
		return cartResponses, err
	}
	cartResponse := dtos.CartResponse{
		CartID:    cart.ID,
		UserID:    cart.UserID,
		ProductID: cart.ProductID,
		Quantity:  cart.Quantity,
		Price:     cart.Price,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}
	return cartResponse, nil
}

// CreateCart godoc
// @Summary      Create a new cart
// @Description  Create a new cart
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Param        request body dtos.CartInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.CartStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /cart [post]
// @Security BearerAuth
func (u *cartUsecase) CreateCart(cart *dtos.CartInput) (dtos.CartResponse, error) {
	var cartResponses dtos.CartResponse

	createCart := models.Cart{
		UserID:    cart.UserID,
		ProductID: cart.ProductID,
		Quantity:  cart.Quantity,
		Price:     cart.Price,
	}

	createdCart, err := u.cartRepo.CreateCart(createCart)
	if err != nil {
		return cartResponses, err
	}

	cartResponse := dtos.CartResponse{
		CartID:    createdCart.ID,
		UserID:    cart.UserID,
		ProductID: cart.ProductID,
		Quantity:  cart.Quantity,
		Price:     cart.Price,
		CreatedAt: createdCart.CreatedAt,
		UpdatedAt: createdCart.UpdatedAt,
	}

	return cartResponse, nil
}

// UpdateCart godoc
// @Summary      Update cart
// @Description  Update cart
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Param id path integer true "ID cart"
// @Param        request body dtos.CartInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.CartStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /cart/{id} [put]
// @Security BearerAuth
func (u *cartUsecase) UpdateCart(id uint, cartInput dtos.CartInput) (dtos.CartResponse, error) {

	var cart models.Cart
	var cartResponse dtos.CartResponse

	cart, err := u.cartRepo.GetCartByID(id)
	if err != nil {
		return cartResponse, err
	}

	cart.ID = id
	cart.UserID = cartInput.UserID
	cart.ProductID = cartInput.ProductID
	cart.Quantity = cartInput.Quantity
	cart.Price = cartInput.Price

	cart, err = u.cartRepo.UpdateCart(cart)

	if err != nil {
		return cartResponse, err
	}

	cartResponse.CartID = cart.ID
	cartResponse.UserID = cart.UserID
	cartResponse.ProductID = cart.ProductID
	cartResponse.Quantity = cart.Quantity
	cartResponse.Price = cart.Price
	cartResponse.CreatedAt = cart.CreatedAt
	cartResponse.UpdatedAt = cart.UpdatedAt

	return cartResponse, nil

}

// DeleteCart godoc
// @Summary      Delete a cart
// @Description  Delete a cart
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Param id path integer true "ID cart"
// @Success      200 {object} dtos.StatusOKDeletedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /cart/{id} [delete]
// @Security BearerAuth
func (u *cartUsecase) DeleteCart(id uint) error {
	cart, err := u.cartRepo.GetCartByID(id)

	if err != nil {
		return nil
	}
	err = u.cartRepo.DeleteCart(cart)
	return err
}

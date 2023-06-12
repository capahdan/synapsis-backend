package usecases

import (
	"synapsis-backend/dtos"
	"synapsis-backend/models"
	"synapsis-backend/repositories"
)

type OrderUsecase interface {
	GetAllOrders(page, limit, user_id int) ([]dtos.OrderResponse, int, error)
	GetOrderByID(id uint) (dtos.OrderResponse, error)
	CreateOrder(order *dtos.OrderInput) (dtos.OrderResponse, error)
	Checkout(order *dtos.OrderInputCheckout) (dtos.OrderResponseCheckout, error)
	UpdateOrder(id uint, orderInput dtos.OrderInput) (dtos.OrderResponse, error)
	DeleteOrder(id uint) error
}

type orderUsecase struct {
	orderRepo       repositories.OrderRepository
	cartRepo        repositories.CartRepository
	productRepo     repositories.ProductRepository
	orderDetailRepo repositories.OrderDetailRepository
}

func NewOrderUsecase(
	OrderRepo repositories.OrderRepository,
	CartRepo repositories.CartRepository,
	ProdutRepo repositories.ProductRepository,
	OrderDetailRepo repositories.OrderDetailRepository,

) OrderUsecase {
	return &orderUsecase{OrderRepo, CartRepo, ProdutRepo, OrderDetailRepo}
}

// GetAllOrders godoc
// @Summary      Get all order
// @Description  Get all order
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param user_id query int false "Seacrh by category ID"
// @Success      200 {object} dtos.GetAllOrderStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /order [get]
// @Security BearerAuth
func (u *orderUsecase) GetAllOrders(page, limit, user_id int) ([]dtos.OrderResponse, int, error) {
	orders, count, err := u.orderRepo.GetAllOrders(page, limit, user_id)
	if err != nil {
		return nil, 0, err
	}

	var orderResponses []dtos.OrderResponse
	for _, order := range orders {
		// category, err := u.orderRepo.GetCategoryByID(order.CategoryID)

		orderResponse := dtos.OrderResponse{
			OrderID:    order.ID,
			TotalPrice: order.TotalPrice,
			UserID:     order.UserID,
			Status:     order.Status,
			CreatedAt:  order.CreatedAt,
			UpdatedAt:  order.UpdatedAt,
		}
		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses, count, nil
}

// GetOrderByID godoc
// @Summary      Get order by ID
// @Description  Get order by ID
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param id path integer true "ID order"
// @Success      200 {object} dtos.OrderStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /order/{id} [get]
// @Security BearerAuth
func (u *orderUsecase) GetOrderByID(id uint) (dtos.OrderResponse, error) {
	var orderResponses dtos.OrderResponse
	order, err := u.orderRepo.GetOrderByID(id)
	if err != nil {
		return orderResponses, err
	}
	orderResponse := dtos.OrderResponse{
		OrderID:    order.ID,
		TotalPrice: order.TotalPrice,
		UserID:     order.UserID,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}
	return orderResponse, nil
}

// CreateOrder godoc
// @Summary      Create a new order
// @Description  Create a new order
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        request body dtos.OrderInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.OrderStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /order [post]
// @Security BearerAuth
func (u *orderUsecase) CreateOrder(order *dtos.OrderInput) (dtos.OrderResponse, error) {
	var orderResponses dtos.OrderResponse

	createOrder := models.Order{
		UserID:     order.UserID,
		TotalPrice: order.TotalPrice,
		Status:     "unpaid",
	}

	createdOrder, err := u.orderRepo.CreateOrder(createOrder)
	if err != nil {
		return orderResponses, err
	}

	orderResponse := dtos.OrderResponse{
		OrderID:    createdOrder.ID,
		TotalPrice: createdOrder.TotalPrice,
		UserID:     createdOrder.UserID,
		Status:     createdOrder.Status,
		CreatedAt:  createdOrder.CreatedAt,
		UpdatedAt:  createdOrder.UpdatedAt,
	}

	return orderResponse, nil
}

// CreateOrder godoc
// @Summary      Create a new order from cart
// @Description  Create a new order from cart
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Param        request body dtos.OrderInput. true "Payload Body [RAW]"
// @Success      200 {object} dtos.OrderCheckoutStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /cart/checkout [post]
// @Security BearerAuth
func (u *orderUsecase) Checkout(order *dtos.OrderInputCheckout) (dtos.OrderResponseCheckout, error) {
	var orderResponses dtos.OrderResponseCheckout
	// First We need to get all carts by user_id
	totalPrice := 0
	page, limit := 1, 100
	carts, _, err := u.cartRepo.GetAllCarts(page, limit, int(order.UserID))
	if err != nil {
		return orderResponses, err
	}

	for _, cart := range carts {
		totalPrice += cart.Price
	}
	// Second We Need to add all price from carts
	// Then We need to create order
	createOrder := models.Order{
		UserID:     order.UserID,
		TotalPrice: totalPrice,
		Status:     "unpaid",
	}

	createdOrder, err := u.orderRepo.CreateOrder(createOrder)
	if err != nil {
		return orderResponses, err
	}

	orderDetailResponses := []dtos.OrderDetailResponse{}
	// Third We need to update stock from product
	// we make record Data in order_detail Table
	for _, cart := range carts {
		product, err := u.productRepo.GetProductByID(cart.ProductID)
		if err != nil {
			return orderResponses, err
		}
		product.Stock -= cart.Quantity
		_, err = u.productRepo.UpdateProduct(product)
		if err != nil {
			return orderResponses, err
		}
		// we Create Order Detail
		createOrderDetail := models.OrderDetail{
			ProductID: product.ID,
			OrderID:   createdOrder.ID,
			Quantity:  cart.Quantity,
			SubTotal:  cart.Price,
		}

		createdOrderDetail, err := u.orderDetailRepo.CreateOrderDetail(createOrderDetail)
		if err != nil {
			return orderResponses, err
		}
		orderDetail := dtos.OrderDetailResponse{
			OrderDetailID: createdOrderDetail.ID,
			ProductID:     createdOrderDetail.ProductID,
			OrderID:       createdOrderDetail.OrderID,
			Quantity:      createdOrderDetail.Quantity,
			SubTotal:      createdOrderDetail.SubTotal,
			CreatedAt:     createdOrderDetail.CreatedAt,
			UpdatedAt:     createdOrderDetail.UpdatedAt,
		}

		orderDetailResponses = append(orderDetailResponses, orderDetail)
		// And delete all carts

	}

	orderResponse := dtos.OrderResponseCheckout{
		OrderID:     createdOrder.ID,
		TotalPrice:  createdOrder.TotalPrice,
		UserID:      createdOrder.UserID,
		Status:      createdOrder.Status,
		OrderDetail: orderDetailResponses,
		CreatedAt:   createdOrder.CreatedAt,
		UpdatedAt:   createdOrder.UpdatedAt,
	}

	return orderResponse, nil
}

// UpdateOrder godoc
// @Summary      Update order
// @Description  Update order
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param id path integer true "ID order"
// @Param        request body dtos.OrderInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.OrderStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /order/{id} [put]
// @Security BearerAuth
func (u *orderUsecase) UpdateOrder(id uint, orderInput dtos.OrderInput) (dtos.OrderResponse, error) {

	var order models.Order
	var orderResponse dtos.OrderResponse

	order, err := u.orderRepo.GetOrderByID(id)
	if err != nil {
		return orderResponse, err
	}

	order.ID = id
	order.UserID = orderInput.UserID
	order.TotalPrice = orderInput.TotalPrice
	// order.Status = orderInput.Status
	order, err = u.orderRepo.UpdateOrder(order)

	if err != nil {
		return orderResponse, err
	}

	orderResponse.OrderID = order.ID
	orderResponse.UserID = order.UserID
	orderResponse.TotalPrice = order.TotalPrice
	orderResponse.Status = order.Status
	orderResponse.CreatedAt = order.CreatedAt
	orderResponse.UpdatedAt = order.UpdatedAt

	return orderResponse, nil

}

// DeleteOrder godoc
// @Summary      Delete a order
// @Description  Delete a order
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param id path integer true "ID order"
// @Success      200 {object} dtos.StatusOKDeletedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /order/{id} [delete]
// @Security BearerAuth
func (u *orderUsecase) DeleteOrder(id uint) error {
	order, err := u.orderRepo.GetOrderByID(id)

	if err != nil {
		return nil
	}
	err = u.orderRepo.DeleteOrder(order)
	return err
}

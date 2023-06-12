package usecases

import (
	"synapsis-backend/dtos"
	"synapsis-backend/models"
	"synapsis-backend/repositories"
)

type OrderDetailUsecase interface {
	GetAllOrderDetails(page, limit, user_id int) ([]dtos.OrderDetailResponse, int, error)
	GetOrderDetailByID(id uint) (dtos.OrderDetailResponse, error)
	CreateOrderDetail(orderDetail *dtos.OrderDetailInput) (dtos.OrderDetailResponse, error)
	UpdateOrderDetail(id uint, orderDetailInput dtos.OrderDetailInput) (dtos.OrderDetailResponse, error)
	DeleteOrderDetail(id uint) error
}

type orderDetailUsecase struct {
	orderDetailRepo repositories.OrderDetailRepository
}

func NewOrderDetailUsecase(OrderDetailRepo repositories.OrderDetailRepository) OrderDetailUsecase {
	return &orderDetailUsecase{OrderDetailRepo}
}

// GetAllOrderDetails godoc
// @Summary      Get all orderDetail
// @Description  Get all orderDetail
// @Tags         OrderDetail
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param user_id query int false "Seacrh by category ID"
// @Success      200 {object} dtos.GetAllOrderDetailStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /orderDetail [get]
// @Security BearerAuth
func (u *orderDetailUsecase) GetAllOrderDetails(page, limit, user_id int) ([]dtos.OrderDetailResponse, int, error) {
	orderDetails, count, err := u.orderDetailRepo.GetAllOrderDetails(page, limit, user_id)
	if err != nil {
		return nil, 0, err
	}

	var orderDetailResponses []dtos.OrderDetailResponse
	for _, orderDetail := range orderDetails {
		// category, err := u.orderDetailRepo.GetCategoryByID(orderDetail.CategoryID)

		orderDetailResponse := dtos.OrderDetailResponse{
			OrderDetailID: orderDetail.ID,
			ProductID:     orderDetail.ProductID,
			OrderID:       orderDetail.OrderID,
			Quantity:      orderDetail.Quantity,
			SubTotal:      orderDetail.SubTotal,
			Discount:      orderDetail.Discount,
			CreatedAt:     orderDetail.CreatedAt,
			UpdatedAt:     orderDetail.UpdatedAt,
		}
		orderDetailResponses = append(orderDetailResponses, orderDetailResponse)
	}

	return orderDetailResponses, count, nil
}

// GetOrderDetailByID godoc
// @Summary      Get orderDetail by ID
// @Description  Get orderDetail by ID
// @Tags         OrderDetail
// @Accept       json
// @Produce      json
// @Param id path integer true "ID orderDetail"
// @Success      200 {object} dtos.OrderDetailStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /orderDetail/{id} [get]
// @Security BearerAuth
func (u *orderDetailUsecase) GetOrderDetailByID(id uint) (dtos.OrderDetailResponse, error) {
	var orderDetailResponses dtos.OrderDetailResponse
	orderDetail, err := u.orderDetailRepo.GetOrderDetailByID(id)
	if err != nil {
		return orderDetailResponses, err
	}
	orderDetailResponse := dtos.OrderDetailResponse{
		OrderDetailID: orderDetail.ID,
		ProductID:     orderDetail.ProductID,
		OrderID:       orderDetail.OrderID,
		Quantity:      orderDetail.Quantity,
		SubTotal:      orderDetail.SubTotal,
		Discount:      orderDetail.Discount,
		CreatedAt:     orderDetail.CreatedAt,
		UpdatedAt:     orderDetail.UpdatedAt,
	}
	return orderDetailResponse, nil
}

// CreateOrderDetail godoc
// @Summary      Create a new orderDetail
// @Description  Create a new orderDetail
// @Tags         OrderDetail
// @Accept       json
// @Produce      json
// @Param        request body dtos.OrderDetailInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.OrderDetailStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /orderDetail [post]
// @Security BearerAuth
func (u *orderDetailUsecase) CreateOrderDetail(orderDetail *dtos.OrderDetailInput) (dtos.OrderDetailResponse, error) {
	var orderDetailResponses dtos.OrderDetailResponse

	createOrderDetail := models.OrderDetail{
		ProductID: orderDetail.ProductID,
		OrderID:   orderDetail.OrderID,
		Quantity:  orderDetail.Quantity,
		SubTotal:  orderDetail.SubTotal,
		Discount:  orderDetail.Discount,
	}

	createdOrderDetail, err := u.orderDetailRepo.CreateOrderDetail(createOrderDetail)
	if err != nil {
		return orderDetailResponses, err
	}

	orderDetailResponse := dtos.OrderDetailResponse{
		OrderDetailID: createdOrderDetail.ID,
		ProductID:     orderDetail.ProductID,
		OrderID:       orderDetail.OrderID,
		Quantity:      orderDetail.Quantity,
		SubTotal:      orderDetail.SubTotal,
		Discount:      orderDetail.Discount,
		CreatedAt:     createdOrderDetail.CreatedAt,
		UpdatedAt:     createdOrderDetail.UpdatedAt,
	}

	return orderDetailResponse, nil
}

// UpdateOrderDetail godoc
// @Summary      Update orderDetail
// @Description  Update orderDetail
// @Tags         OrderDetail
// @Accept       json
// @Produce      json
// @Param id path integer true "ID orderDetail"
// @Param        request body dtos.OrderDetailInput true "Payload Body [RAW]"
// @Success      200 {object} dtos.OrderDetailStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /orderDetail/{id} [put]
// @Security BearerAuth
func (u *orderDetailUsecase) UpdateOrderDetail(id uint, orderDetailInput dtos.OrderDetailInput) (dtos.OrderDetailResponse, error) {

	var orderDetail models.OrderDetail
	var orderDetailResponse dtos.OrderDetailResponse

	orderDetail, err := u.orderDetailRepo.GetOrderDetailByID(id)
	if err != nil {
		return orderDetailResponse, err
	}

	orderDetail.ID = id
	orderDetail.ProductID = orderDetailInput.ProductID
	orderDetail.OrderID = orderDetailInput.OrderID
	orderDetail.Quantity = orderDetailInput.Quantity
	orderDetail.SubTotal = orderDetailInput.SubTotal
	orderDetail.Discount = orderDetailInput.Discount

	orderDetail, err = u.orderDetailRepo.UpdateOrderDetail(orderDetail)

	if err != nil {
		return orderDetailResponse, err
	}

	orderDetailResponse.OrderDetailID = orderDetail.ID
	orderDetailResponse.ProductID = orderDetail.ProductID
	orderDetailResponse.OrderID = orderDetail.OrderID
	orderDetailResponse.Quantity = orderDetail.Quantity
	orderDetailResponse.SubTotal = orderDetail.SubTotal
	orderDetailResponse.Discount = orderDetail.Discount
	orderDetailResponse.CreatedAt = orderDetail.CreatedAt
	orderDetailResponse.UpdatedAt = orderDetail.UpdatedAt

	return orderDetailResponse, nil

}

// DeleteOrderDetail godoc
// @Summary      Delete a orderDetail
// @Description  Delete a orderDetail
// @Tags         OrderDetail
// @Accept       json
// @Produce      json
// @Param id path integer true "ID orderDetail"
// @Success      200 {object} dtos.StatusOKDeletedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /orderDetail/{id} [delete]
// @Security BearerAuth
func (u *orderDetailUsecase) DeleteOrderDetail(id uint) error {
	orderDetail, err := u.orderDetailRepo.GetOrderDetailByID(id)

	if err != nil {
		return nil
	}
	err = u.orderDetailRepo.DeleteOrderDetail(orderDetail)
	return err
}

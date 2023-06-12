package repositories

import (
	"synapsis-backend/models"

	"gorm.io/gorm"
)

type OrderDetailRepository interface {
	GetAllOrderDetails(page, limit, user_id int) ([]models.OrderDetail, int, error)
	GetOrderDetailByID(id uint) (models.OrderDetail, error)
	CreateOrderDetail(orderDetail models.OrderDetail) (models.OrderDetail, error)
	UpdateOrderDetail(orderDetail models.OrderDetail) (models.OrderDetail, error)
	DeleteOrderDetail(orderDetail models.OrderDetail) error
}

type orderDetailRepository struct {
	db *gorm.DB
}

func NewOrderDetailRepository(db *gorm.DB) OrderDetailRepository {
	return &orderDetailRepository{db}
}

// Implementasi fungsi-fungsi dari interface ItemRepository

func (r *orderDetailRepository) GetAllOrderDetails(page, limit, user_id int) ([]models.OrderDetail, int, error) {
	var (
		orderDetails []models.OrderDetail
		count        int64
	)
	err := r.db.Find(&orderDetails).Count(&count).Error
	if err != nil {
		return orderDetails, int(count), err
	}

	offset := (page - 1) * limit

	if user_id != 0 {
		err = r.db.Where("user_id = ?", user_id).Limit(limit).Offset(offset).Find(&orderDetails).Count(&count).Error
		return orderDetails, int(count), err
	}
	err = r.db.Limit(limit).Offset(offset).Find(&orderDetails).Error

	return orderDetails, int(count), err
}

func (r *orderDetailRepository) GetOrderDetailByID(id uint) (models.OrderDetail, error) {
	var orderDetail models.OrderDetail
	err := r.db.Where("id = ?", id).First(&orderDetail).Error
	return orderDetail, err
}

func (r *orderDetailRepository) CreateOrderDetail(orderDetail models.OrderDetail) (models.OrderDetail, error) {
	err := r.db.Create(&orderDetail).Error
	return orderDetail, err
}

func (r *orderDetailRepository) UpdateOrderDetail(orderDetail models.OrderDetail) (models.OrderDetail, error) {
	err := r.db.Save(&orderDetail).Error
	return orderDetail, err
}

func (r *orderDetailRepository) DeleteOrderDetail(orderDetail models.OrderDetail) error {
	err := r.db.Delete(&orderDetail).Error
	return err
}

package repositories

import (
	"synapsis-backend/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAllOrders(page, limit, user_id int) ([]models.Order, int, error)
	GetOrderByID(id uint) (models.Order, error)
	CreateOrder(order models.Order) (models.Order, error)
	UpdateOrder(order models.Order) (models.Order, error)
	DeleteOrder(order models.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

// Implementasi fungsi-fungsi dari interface ItemRepository

func (r *orderRepository) GetAllOrders(page, limit, user_id int) ([]models.Order, int, error) {
	var (
		orders []models.Order
		count  int64
	)
	err := r.db.Find(&orders).Count(&count).Error
	if err != nil {
		return orders, int(count), err
	}

	offset := (page - 1) * limit

	if user_id != 0 {
		err = r.db.Where("user_id = ?", user_id).Limit(limit).Offset(offset).Find(&orders).Count(&count).Error
		return orders, int(count), err
	}
	err = r.db.Limit(limit).Offset(offset).Find(&orders).Error

	return orders, int(count), err
}

func (r *orderRepository) GetOrderByID(id uint) (models.Order, error) {
	var order models.Order
	err := r.db.Where("id = ?", id).First(&order).Error
	return order, err
}

func (r *orderRepository) CreateOrder(order models.Order) (models.Order, error) {
	err := r.db.Create(&order).Error
	return order, err
}

func (r *orderRepository) UpdateOrder(order models.Order) (models.Order, error) {
	err := r.db.Save(&order).Error
	return order, err
}

func (r *orderRepository) DeleteOrder(order models.Order) error {
	err := r.db.Delete(&order).Error
	return err
}

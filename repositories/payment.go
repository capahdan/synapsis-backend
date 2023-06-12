package repositories

import (
	"synapsis-backend/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	GetAllPayments(page, limit, category_id int) ([]models.Payment, int, error)
	GetPaymentByID(id uint) (models.Payment, error)
	CreatePayment(payment models.Payment) (models.Payment, error)
	UpdatePayment(payment models.Payment) (models.Payment, error)
	DeletePayment(payment models.Payment) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

// Implementasi fungsi-fungsi dari interface ItemRepository

func (r *paymentRepository) GetAllPayments(page, limit, category_id int) ([]models.Payment, int, error) {
	var (
		payments []models.Payment
		count    int64
	)
	err := r.db.Find(&payments).Count(&count).Error
	if err != nil {
		return payments, int(count), err
	}

	offset := (page - 1) * limit

	if category_id != 0 {
		err = r.db.Where("category_id = ?", category_id).Limit(limit).Offset(offset).Find(&payments).Count(&count).Error
		return payments, int(count), err
	}
	err = r.db.Limit(limit).Offset(offset).Find(&payments).Error

	return payments, int(count), err
}

func (r *paymentRepository) GetPaymentByID(id uint) (models.Payment, error) {
	var payment models.Payment
	err := r.db.Where("id = ?", id).First(&payment).Error
	return payment, err
}

func (r *paymentRepository) CreatePayment(payment models.Payment) (models.Payment, error) {
	err := r.db.Create(&payment).Error
	return payment, err
}

func (r *paymentRepository) UpdatePayment(payment models.Payment) (models.Payment, error) {
	err := r.db.Save(&payment).Error
	return payment, err
}

func (r *paymentRepository) DeletePayment(payment models.Payment) error {
	err := r.db.Delete(&payment).Error
	return err
}

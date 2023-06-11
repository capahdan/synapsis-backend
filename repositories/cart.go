package repositories

import (
	"synapsis-backend/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	GetAllCarts(page, limit, user_id int) ([]models.Cart, int, error)
	GetCartByID(id uint) (models.Cart, error)
	CreateCart(cart models.Cart) (models.Cart, error)
	UpdateCart(cart models.Cart) (models.Cart, error)
	DeleteCart(cart models.Cart) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

// Implementasi fungsi-fungsi dari interface ItemRepository

func (r *cartRepository) GetAllCarts(page, limit, user_id int) ([]models.Cart, int, error) {
	var (
		carts []models.Cart
		count int64
	)
	err := r.db.Find(&carts).Count(&count).Error
	if err != nil {
		return carts, int(count), err
	}

	offset := (page - 1) * limit

	if user_id != 0 {
		err = r.db.Where("user_id = ?", user_id).Limit(limit).Offset(offset).Find(&carts).Count(&count).Error
		return carts, int(count), err
	}
	err = r.db.Limit(limit).Offset(offset).Find(&carts).Error

	return carts, int(count), err
}

func (r *cartRepository) GetCartByID(id uint) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Where("id = ?", id).First(&cart).Error
	return cart, err
}

func (r *cartRepository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error
	return cart, err
}

func (r *cartRepository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Save(&cart).Error
	return cart, err
}

func (r *cartRepository) DeleteCart(cart models.Cart) error {
	err := r.db.Delete(&cart).Error
	return err
}

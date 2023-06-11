package repositories

import (
	"synapsis-backend/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts(page, limit, category_id int) ([]models.Product, int, error)
	GetProductByID(id uint) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(product models.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

// Implementasi fungsi-fungsi dari interface ItemRepository

func (r *productRepository) GetAllProducts(page, limit, category_id int) ([]models.Product, int, error) {
	var (
		products []models.Product
		count    int64
	)
	err := r.db.Find(&products).Count(&count).Error
	if err != nil {
		return products, int(count), err
	}

	offset := (page - 1) * limit

	if category_id != 0 {
		err = r.db.Where("category_id = ?", category_id).Limit(limit).Offset(offset).Find(&products).Count(&count).Error
		return products, int(count), err
	}
	err = r.db.Limit(limit).Offset(offset).Find(&products).Error

	return products, int(count), err
}

func (r *productRepository) GetProductByID(id uint) (models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ?", id).First(&product).Error
	return product, err
}

func (r *productRepository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *productRepository) UpdateProduct(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

func (r *productRepository) DeleteProduct(product models.Product) error {
	err := r.db.Delete(&product).Error
	return err
}

package repositories

import (
	"synapsis-backend/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategorys(page, limit int) ([]models.Category, int, error)
	GetCategoryByID(id uint) (models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
	UpdateCategory(category models.Category) (models.Category, error)
	DeleteCategory(category models.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

// Implementasi fungsi-fungsi dari interface ItemRepository

func (r *categoryRepository) GetAllCategorys(page, limit int) ([]models.Category, int, error) {
	var (
		categorys []models.Category
		count     int64
	)
	err := r.db.Find(&categorys).Count(&count).Error
	if err != nil {
		return categorys, int(count), err
	}

	offset := (page - 1) * limit

	err = r.db.Limit(limit).Offset(offset).Find(&categorys).Error

	return categorys, int(count), err
}

func (r *categoryRepository) GetCategoryByID(id uint) (models.Category, error) {
	var category models.Category
	err := r.db.Where("id = ?", id).First(&category).Error
	return category, err
}

func (r *categoryRepository) CreateCategory(category models.Category) (models.Category, error) {
	err := r.db.Create(&category).Error
	return category, err
}

func (r *categoryRepository) UpdateCategory(category models.Category) (models.Category, error) {
	err := r.db.Save(&category).Error
	return category, err
}

func (r *categoryRepository) DeleteCategory(category models.Category) error {
	err := r.db.Delete(&category).Error
	return err
}

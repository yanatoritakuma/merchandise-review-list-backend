package repository

import (
	"merchandise-review-list-backend/model"

	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(product *model.Product) error
	GetMyProducts(product *[]model.Product, userId uint, page int, pageSize int) (int, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{db}
}

func (pr *productRepository) CreateProduct(product *model.Product) error {
	if err := pr.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (pr *productRepository) GetMyProducts(product *[]model.Product, userId uint, page int, pageSize int) (int, error) {
	offset := (page - 1) * pageSize
	var totalCount int64

	if err := pr.db.Model(&model.Product{}).Where("user_id=?", userId).Count(&totalCount).Error; err != nil {
		return 0, err
	}

	if err := pr.db.Joins("User").Where("user_id=?", userId).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(product).Error; err != nil {
		return 0, err
	}
	return int(totalCount), nil
}

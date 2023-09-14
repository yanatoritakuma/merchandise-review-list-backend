package repository

import (
	"merchandise-review-list-backend/model"

	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(product *model.Product) error
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

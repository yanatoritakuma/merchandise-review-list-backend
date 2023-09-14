package usecase

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/repository"
)

type IProductUsecase interface {
	CreateProduct(product model.Product) (model.ProductResponse, error)
}

type productUsecase struct {
	pr repository.IProductRepository
}

func NweProductUsecase(pr repository.IProductRepository) IProductUsecase {
	return &productUsecase{pr}
}

func (pu *productUsecase) CreateProduct(product model.Product) (model.ProductResponse, error) {
	if err := pu.pr.CreateProduct(&product); err != nil {
		return model.ProductResponse{}, err
	}
	resProduct := model.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
		Price:       product.Price,
		Review:      product.Review,
		Url:         product.Url,
		Image:       product.Image,
		CreatedAt:   product.CreatedAt,
	}
	return resProduct, nil
}

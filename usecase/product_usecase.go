package usecase

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/repository"
	"merchandise-review-list-backend/validator"
	"time"
)

type IProductUsecase interface {
	CreateProduct(product model.Product) (model.ProductResponse, error)
	UpdateTimeLimit(product model.Product, userId uint, productId uint) (model.ProductResponse, error)
	DeleteProduct(userId uint, productId uint) error
	GetMyProducts(userId uint, page int, pageSize int) ([]model.ProductResponse, int, error)
	GetMyProductsTimeLimitAll(userId uint, page int, pageSize int) ([]model.ProductResponse, int, error)
	GetMyProductsTimeLimitYearMonth(userId uint, yearMonth time.Time) ([]model.ProductYearMonthResponse, error)
	GetMyProductsTimeLimitDate(userId uint, page int, pageSize int, date time.Time) ([]model.ProductResponse, int, error)
}

type productUsecase struct {
	pr repository.IProductRepository
	pv validator.IProductValidator
}

func NweProductUsecase(pr repository.IProductRepository, pv validator.IProductValidator) IProductUsecase {
	return &productUsecase{pr, pv}
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
		Provider:    product.Provider,
		TimeLimit:   product.TimeLimit,
		CreatedAt:   product.CreatedAt,
	}
	return resProduct, nil
}

func (pu *productUsecase) UpdateTimeLimit(product model.Product, userId uint, productId uint) (model.ProductResponse, error) {
	if err := pu.pv.ProductValidator(product); err != nil {
		return model.ProductResponse{}, err
	}
	if err := pu.pr.UpdateTimeLimit(&product, userId, productId); err != nil {
		return model.ProductResponse{}, err
	}

	resProduct := model.ProductResponse{
		TimeLimit: product.TimeLimit,
	}
	return resProduct, nil
}

func (pu *productUsecase) DeleteProduct(userId uint, productId uint) error {
	if err := pu.pr.DeleteProduct(userId, productId); err != nil {
		return err
	}
	return nil
}

func (pu *productUsecase) GetMyProducts(userId uint, page int, pageSize int) ([]model.ProductResponse, int, error) {
	product := []model.Product{}

	totalCount, err := pu.pr.GetMyProducts(&product, userId, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	resProducts := []model.ProductResponse{}
	for _, product := range product {

		p := model.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Stock:       product.Stock,
			Price:       product.Price,
			Review:      product.Review,
			Url:         product.Url,
			Image:       product.Image,
			Provider:    product.Provider,
			TimeLimit:   product.TimeLimit,
			CreatedAt:   product.CreatedAt,
		}
		resProducts = append(resProducts, p)
	}

	return resProducts, totalCount, nil
}

func (pu *productUsecase) GetMyProductsTimeLimitAll(userId uint, page int, pageSize int) ([]model.ProductResponse, int, error) {
	product := []model.Product{}

	totalCount, err := pu.pr.GetMyProductsTimeLimitAll(&product, userId, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	resProducts := []model.ProductResponse{}
	for _, product := range product {

		p := model.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Stock:       product.Stock,
			Price:       product.Price,
			Review:      product.Review,
			Url:         product.Url,
			Image:       product.Image,
			Provider:    product.Provider,
			TimeLimit:   product.TimeLimit,
			CreatedAt:   product.CreatedAt,
		}
		resProducts = append(resProducts, p)
	}

	return resProducts, totalCount, nil
}

func (pu *productUsecase) GetMyProductsTimeLimitYearMonth(userId uint, yearMonth time.Time) ([]model.ProductYearMonthResponse, error) {
	product := []model.Product{}

	err := pu.pr.GetMyProductsTimeLimitYearMonth(&product, userId, yearMonth)
	if err != nil {
		return nil, err
	}

	resProducts := []model.ProductYearMonthResponse{}
	for _, product := range product {

		p := model.ProductYearMonthResponse{
			TimeLimit: product.TimeLimit,
		}
		resProducts = append(resProducts, p)
	}

	return resProducts, nil
}

func (pu *productUsecase) GetMyProductsTimeLimitDate(userId uint, page int, pageSize int, date time.Time) ([]model.ProductResponse, int, error) {
	product := []model.Product{}

	totalCount, err := pu.pr.GetMyProductsTimeLimitDate(&product, userId, page, pageSize, date)
	if err != nil {
		return nil, 0, err
	}

	resProducts := []model.ProductResponse{}
	for _, product := range product {

		p := model.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Stock:       product.Stock,
			Price:       product.Price,
			Review:      product.Review,
			Url:         product.Url,
			Image:       product.Image,
			Provider:    product.Provider,
			TimeLimit:   product.TimeLimit,
			CreatedAt:   product.CreatedAt,
		}
		resProducts = append(resProducts, p)
	}

	return resProducts, totalCount, nil
}

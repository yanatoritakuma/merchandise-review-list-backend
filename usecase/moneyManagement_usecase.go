package usecase

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/repository"
	"time"
)

type IMoneyManagementUsecase interface {
	CreateMoneyManagement(moneyManagement model.MoneyManagement) (model.MoneyManagementResponse, error)
	GetMyMoneyManagements(userId uint, yearMonth time.Time) (model.MoneyManagementByCategoryResponse, error)
}

type moneyManagementUsecase struct {
	mr repository.IMoneyManagementRepository
}

func NewMoneyManagementUsecase(
	mr repository.IMoneyManagementRepository,
) IMoneyManagementUsecase {
	return &moneyManagementUsecase{mr}
}

func (mu *moneyManagementUsecase) CreateMoneyManagement(moneyManagement model.MoneyManagement) (model.MoneyManagementResponse, error) {
	// todo:バリデーション未実装

	if err := mu.mr.CreateMoneyManagement(&moneyManagement); err != nil {
		return model.MoneyManagementResponse{}, err
	}

	resMoneyManagement := model.MoneyManagementResponse{
		ID:         moneyManagement.ID,
		Title:      moneyManagement.Title,
		Category:   moneyManagement.Category,
		UnitPrice:  moneyManagement.UnitPrice,
		Quantity:   moneyManagement.Quantity,
		TotalPrice: moneyManagement.TotalPrice,
		CreatedAt:  moneyManagement.CreatedAt,
		UpdatedAt:  moneyManagement.UpdatedAt,
	}
	return resMoneyManagement, nil
}

func (mu *moneyManagementUsecase) GetMyMoneyManagements(userId uint, yearMonth time.Time) (model.MoneyManagementByCategoryResponse, error) {
	moneyManagement := []model.MoneyManagement{}
	err := mu.mr.GetMyMoneyManagements(&moneyManagement, userId, yearMonth)
	if err != nil {
		return model.MoneyManagementByCategoryResponse{}, err
	}

	res := model.MoneyManagementByCategoryResponse{
		Food:          []model.MoneyManagementResponse{},
		Drink:         []model.MoneyManagementResponse{},
		Book:          []model.MoneyManagementResponse{},
		Fashion:       []model.MoneyManagementResponse{},
		Furniture:     []model.MoneyManagementResponse{},
		GamesToys:     []model.MoneyManagementResponse{},
		Beauty:        []model.MoneyManagementResponse{},
		EveryDayItems: []model.MoneyManagementResponse{},
		Other:         []model.MoneyManagementResponse{},
	}

	for _, mm := range moneyManagement {
		switch mm.Category {
		case "food":
			res.Food = append(res.Food, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
			})
		case "drink":
			res.Drink = append(res.Drink, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
			})
		case "book":
			res.Book = append(res.Book, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
			})
		case "fashion":
			res.Fashion = append(res.Fashion, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
			})
		case "furniture":
			res.Furniture = append(res.Furniture, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
			})
		case "gamesToys":
			res.GamesToys = append(res.GamesToys, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
			})
		case "beauty":
			res.Beauty = append(res.Beauty, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
			})
		case "everyDayItems":
			res.EveryDayItems = append(res.EveryDayItems, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
			})
		case "other":
			res.Other = append(res.Other, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
			})

		}
	}

	return res, nil
}

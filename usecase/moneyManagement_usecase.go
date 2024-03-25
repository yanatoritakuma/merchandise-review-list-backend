package usecase

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/repository"
	"merchandise-review-list-backend/validator"
	"time"
)

type IMoneyManagementUsecase interface {
	CreateMoneyManagement(moneyManagement model.MoneyManagement) (model.MoneyManagementResponse, error)
	GetMyMoneyManagements(userId uint, yearMonth time.Time, yearFlag bool) (model.MoneyManagementByCategoryResponse, error)
}

type moneyManagementUsecase struct {
	mr repository.IMoneyManagementRepository
	mv validator.IMoneyManagementValidator
}

func NewMoneyManagementUsecase(
	mr repository.IMoneyManagementRepository,
	mv validator.IMoneyManagementValidator,
) IMoneyManagementUsecase {
	return &moneyManagementUsecase{mr, mv}
}

func (mu *moneyManagementUsecase) CreateMoneyManagement(moneyManagement model.MoneyManagement) (model.MoneyManagementResponse, error) {
	if err := mu.mv.MoneyManagementValidator(moneyManagement); err != nil {
		return model.MoneyManagementResponse{}, err
	}

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

func (mu *moneyManagementUsecase) GetMyMoneyManagements(userId uint, yearMonth time.Time, yearFlag bool) (model.MoneyManagementByCategoryResponse, error) {
	moneyManagement := []model.MoneyManagement{}
	err := mu.mr.GetMyMoneyManagements(&moneyManagement, userId, yearMonth, yearFlag)
	if err != nil {
		return model.MoneyManagementByCategoryResponse{}, err
	}

	res := model.MoneyManagementByCategoryResponse{
		Food:          model.MoneyManagementByCategoryItemResponse{},
		Drink:         model.MoneyManagementByCategoryItemResponse{},
		Book:          model.MoneyManagementByCategoryItemResponse{},
		Fashion:       model.MoneyManagementByCategoryItemResponse{},
		Furniture:     model.MoneyManagementByCategoryItemResponse{},
		GamesToys:     model.MoneyManagementByCategoryItemResponse{},
		Beauty:        model.MoneyManagementByCategoryItemResponse{},
		EveryDayItems: model.MoneyManagementByCategoryItemResponse{},
		Other:         model.MoneyManagementByCategoryItemResponse{},
	}

	totalPrice := uint(0)

	for _, mm := range moneyManagement {
		switch mm.Category {
		case "food":
			res.Food.Items = append(res.Food.Items, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
				UpdatedAt:  mm.UpdatedAt,
			})

			res.Food.ItemTotalPrice += mm.TotalPrice
			totalPrice += mm.TotalPrice
		case "drink":
			res.Drink.Items = append(res.Drink.Items, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
				UpdatedAt:  mm.UpdatedAt,
			})

			res.Drink.ItemTotalPrice += mm.TotalPrice
			totalPrice += mm.TotalPrice
		case "book":
			res.Book.Items = append(res.Book.Items, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
				UpdatedAt:  mm.UpdatedAt,
			})

			res.Book.ItemTotalPrice += mm.TotalPrice
			totalPrice += mm.TotalPrice
		case "fashion":
			res.Fashion.Items = append(res.Fashion.Items, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
				UpdatedAt:  mm.UpdatedAt,
			})

			res.Fashion.ItemTotalPrice += mm.TotalPrice
			totalPrice += mm.TotalPrice
		case "furniture":
			res.Furniture.Items = append(res.Furniture.Items, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
				UpdatedAt:  mm.UpdatedAt,
			})

			res.Furniture.ItemTotalPrice += mm.TotalPrice
			totalPrice += mm.TotalPrice
		case "gamesToys":
			res.GamesToys.Items = append(res.GamesToys.Items, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
				UpdatedAt:  mm.UpdatedAt,
			})

			res.GamesToys.ItemTotalPrice += mm.TotalPrice
			totalPrice += mm.TotalPrice
		case "beauty":
			res.Beauty.Items = append(res.Beauty.Items, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
				UpdatedAt:  mm.UpdatedAt,
			})

			res.Beauty.ItemTotalPrice += mm.TotalPrice
			totalPrice += mm.TotalPrice
		case "everyDayItems":
			res.EveryDayItems.Items = append(res.EveryDayItems.Items, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
				UpdatedAt:  mm.UpdatedAt,
			})

			res.EveryDayItems.ItemTotalPrice += mm.TotalPrice
			totalPrice += mm.TotalPrice
		case "other":
			res.Other.Items = append(res.Other.Items, model.MoneyManagementResponse{
				ID:         mm.ID,
				Title:      mm.Title,
				Category:   mm.Category,
				UnitPrice:  mm.UnitPrice,
				Quantity:   mm.Quantity,
				TotalPrice: mm.TotalPrice,
				CreatedAt:  mm.CreatedAt,
				UpdatedAt:  mm.UpdatedAt,
			})

			res.Other.ItemTotalPrice += mm.TotalPrice
			totalPrice += mm.TotalPrice
		}

		res.TotalPrice = totalPrice
	}

	return res, nil
}

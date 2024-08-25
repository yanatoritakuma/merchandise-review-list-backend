package usecase

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/repository"
)

type IHouseholdBudgetEstimateItemUsecase interface {
	CreateHouseholdBudgetEstimateItem(householdBudgetEstimateItem model.HouseholdBudgetEstimateItem) (model.HouseholdBudgetEstimateItemResponse, error)
}

type householdBudgetEstimateItemUsecase struct {
	hr repository.IHouseholdBudgetEstimateItemRepository
}

func NweHouseholdBudgetEstimateItemUsecase(hr repository.IHouseholdBudgetEstimateItemRepository) IHouseholdBudgetEstimateItemUsecase {
	return &householdBudgetEstimateItemUsecase{hr}
}

// 家計簿のアイテム作成
func (hu *householdBudgetEstimateItemUsecase) CreateHouseholdBudgetEstimateItem(householdBudgetEstimateItem model.HouseholdBudgetEstimateItem) (model.HouseholdBudgetEstimateItemResponse, error) {
	if err := hu.hr.CreateHouseholdBudgetEstimateItem(&householdBudgetEstimateItem); err != nil {
		return model.HouseholdBudgetEstimateItemResponse{}, err
	}

	resHouseholdBudgetEstimateItem := model.HouseholdBudgetEstimateItemResponse{
		ID:        householdBudgetEstimateItem.ID,
		Name:      householdBudgetEstimateItem.Name,
		CreatedAt: householdBudgetEstimateItem.CreatedAt,
	}

	return resHouseholdBudgetEstimateItem, nil

}

package usecase

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/repository"
)

type IHouseholdBudgetEstimateItemUsecase interface {
	CreateHouseholdBudgetEstimateItem(householdBudgetEstimateItem model.HouseholdBudgetEstimateItem) (model.HouseholdBudgetEstimateItemResponse, error)
	GetMyHouseholdBudgetEstimateItem(householdBudgetId uint, userId uint, year string, month string) ([]model.HouseholdBudgetEstimateItemResponse, error)
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
		CreatedAt: householdBudgetEstimateItem.CreatedAt,
	}

	return resHouseholdBudgetEstimateItem, nil

}

// 特定の家計簿予算の年月日に紐づくアイテムを取得する
func (hu *householdBudgetEstimateItemUsecase) GetMyHouseholdBudgetEstimateItem(householdBudgetId uint, userId uint, year string, month string) ([]model.HouseholdBudgetEstimateItemResponse, error) {
	householdBudgetEstimateItem := []model.HouseholdBudgetEstimateItem{}
	err := hu.hr.GetMyHouseholdBudgetEstimateItem(&householdBudgetEstimateItem, householdBudgetId, userId, year, month)
	if err != nil {
		return nil, err
	}

	resHouseholdBudge := []model.HouseholdBudgetEstimateItemResponse{}

	for _, v := range householdBudgetEstimateItem {
		h := model.HouseholdBudgetEstimateItemResponse{
			ID:         v.ID,
			CategoryId: v.CategoryId,
			CreatedAt:  v.CreatedAt,
		}

		resHouseholdBudge = append(resHouseholdBudge, h)
	}

	return resHouseholdBudge, nil
}

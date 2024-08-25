package repository

import (
	"merchandise-review-list-backend/model"

	"gorm.io/gorm"
)

type IHouseholdBudgetEstimateItemRepository interface {
	CreateHouseholdBudgetEstimateItem(householdBudgetEstimateItem *model.HouseholdBudgetEstimateItem) error
}

type householdBudgetEstimateItemRepository struct {
	db *gorm.DB
}

func NewHouseholdBudgetEstimateItemRepository(db *gorm.DB) IHouseholdBudgetEstimateItemRepository {
	return &householdBudgetEstimateItemRepository{db}
}

// 家計簿予算アイテム作成
func (hr *householdBudgetEstimateItemRepository) CreateHouseholdBudgetEstimateItem(householdBudgetEstimateItem *model.HouseholdBudgetEstimateItem) error {
	if err := hr.db.Create(householdBudgetEstimateItem).Error; err != nil {
		return err
	}
	return nil
}

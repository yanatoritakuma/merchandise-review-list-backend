package repository

import (
	"errors"
	"merchandise-review-list-backend/model"

	"gorm.io/gorm"
)

type IHouseholdBudgetEstimateItemRepository interface {
	CreateHouseholdBudgetEstimateItem(householdBudgetEstimateItem *model.HouseholdBudgetEstimateItem) error
	GetMyHouseholdBudgetEstimateItem(householdBudgetEstimateItem *[]model.HouseholdBudgetEstimateItem, householdBudgetId uint, userId uint, year string, month string) error
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

// 特定の家計簿予算の年月日に紐づくアイテムを取得する
func (hr *householdBudgetEstimateItemRepository) GetMyHouseholdBudgetEstimateItem(householdBudgetEstimateItem *[]model.HouseholdBudgetEstimateItem, householdBudgetId uint, userId uint, year string, month string) error {
	if err := hr.db.Where("household_budget_id = ? AND user_id = ? AND year = ? AND month = ?", householdBudgetId, userId, year, month).Find(householdBudgetEstimateItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return nil
}

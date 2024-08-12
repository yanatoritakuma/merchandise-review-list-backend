package repository

import (
	"merchandise-review-list-backend/model"

	"gorm.io/gorm"
)

type IHouseholdBudgetRepository interface {
	CreateHouseholdBudget(householdBudget *model.HouseholdBudget) error
}

type householdBudgetRepository struct {
	db *gorm.DB
}

func NewHouseholdBudgetRepository(db *gorm.DB) IHouseholdBudgetRepository {
	return &householdBudgetRepository{db}
}

func (hr *householdBudgetRepository) CreateHouseholdBudget(householdBudget *model.HouseholdBudget) error {
	if err := hr.db.Create(householdBudget).Error; err != nil {
		return err
	}
	return nil
}

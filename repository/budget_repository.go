package repository

import (
	"merchandise-review-list-backend/model"

	"gorm.io/gorm"
)

type IBudgetRepository interface {
	CreateBudget(budget *model.Budget) error
}

type budgetRepository struct {
	db *gorm.DB
}

func NewBudgetRepository(db *gorm.DB) IBudgetRepository {
	return &budgetRepository{db}
}

func (br *budgetRepository) CreateBudget(budget *model.Budget) error {
	if err := br.db.Create(budget).Error; err != nil {
		return err
	}
	return nil
}

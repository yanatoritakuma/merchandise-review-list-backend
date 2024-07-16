package repository

import (
	"errors"
	"fmt"
	"merchandise-review-list-backend/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IBudgetRepository interface {
	CreateBudget(budget *model.Budget) error
	UpdateBudget(budget *model.Budget, userId uint, id uint) error
	SameYearMonth(userId uint, year string, month string) (*model.Budget, error) //既に設定年月が存在しているか
	GetBudgetByUserId(budget *model.Budget, userId uint, year string, month string) error
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

func (br *budgetRepository) UpdateBudget(budget *model.Budget, userId uint, id uint) error {
	result := br.db.Model(budget).Clauses(clause.Returning{}).Where("id=? AND user_id=?", id, userId).Updates(map[string]interface{}{
		"month":           budget.Month,
		"year":            budget.Year,
		"total_price":     budget.TotalPrice,
		"food":            budget.Food,
		"drink":           budget.Drink,
		"book":            budget.Book,
		"fashion":         budget.Fashion,
		"furniture":       budget.Furniture,
		"games_toys":      budget.GamesToys,
		"beauty":          budget.Beauty,
		"every_day_items": budget.EveryDayItems,
		"other":           budget.Other,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (br *budgetRepository) SameYearMonth(userId uint, year string, month string) (*model.Budget, error) {
	budget := &model.Budget{}

	if err := br.db.Where("user_id=? AND year=? AND month=?", userId, year, month).First(budget).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 登録済みの設定金額が存在しない場合はnilを返す
			return nil, nil
		}
		return nil, err
	}

	return budget, nil
}

func (br *budgetRepository) GetBudgetByUserId(budget *model.Budget, userId uint, year string, month string) error {
	if err := br.db.Where("user_id=? AND year=? AND month=?", userId, year, month).First(budget).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil // レコードが見つからない場合はnilを返す
		}

		return err
	}
	return nil
}

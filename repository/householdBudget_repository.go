package repository

import (
	"merchandise-review-list-backend/model"

	"gorm.io/gorm"
)

type IHouseholdBudgetRepository interface {
	CreateHouseholdBudget(householdBudget *model.HouseholdBudget) error
	GetMyHouseholdBudget(householdBudget *[]model.HouseholdBudget, userId uint, page int, pageSize int) (int, error)
}

type householdBudgetRepository struct {
	db *gorm.DB
}

func NewHouseholdBudgetRepository(db *gorm.DB) IHouseholdBudgetRepository {
	return &householdBudgetRepository{db}
}

// 家計簿を作成
func (hr *householdBudgetRepository) CreateHouseholdBudget(householdBudget *model.HouseholdBudget) error {
	if err := hr.db.Create(householdBudget).Error; err != nil {
		return err
	}
	return nil
}

// 自分が作成した家計簿取得 todo: 家計簿共有テーブル（household_budget_share）を参照して自分が所属している家計も取得
func (hr *householdBudgetRepository) GetMyHouseholdBudget(householdBudget *[]model.HouseholdBudget, userId uint, page int, pageSize int) (int, error) {
	offset := (page - 1) * pageSize
	var totalCount int64

	if err := hr.db.Where("user_id=?", userId).Model(&model.HouseholdBudget{}).Count(&totalCount).Error; err != nil {
		return 0, err
	}

	if err := hr.db.Where("user_id=?", userId).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(householdBudget).Error; err != nil {
		return 0, err
	}

	return int(totalCount), nil
}

package repository

import (
	"merchandise-review-list-backend/model"
	"time"

	"gorm.io/gorm"
)

type IMoneyManagementRepository interface {
	CreateMoneyManagement(moneyManagement *model.MoneyManagement) error
	GetMyMoneyManagements(moneyManagement *[]model.MoneyManagement, userId uint, yearMonth time.Time) error
}

type moneyManagementRepository struct {
	db *gorm.DB
}

func NewMoneyManagementRepository(db *gorm.DB) IMoneyManagementRepository {
	return &moneyManagementRepository{db}
}

func (mr *moneyManagementRepository) CreateMoneyManagement(moneyManagement *model.MoneyManagement) error {
	if err := mr.db.Create(moneyManagement).Error; err != nil {
		return err
	}
	return nil
}

func (mr *moneyManagementRepository) GetMyMoneyManagements(moneyManagement *[]model.MoneyManagement, userId uint, yearMonth time.Time) error {
	startOfMonth := time.Date(yearMonth.Year(), yearMonth.Month(), 1, 0, 0, 0, 0, time.UTC)

	if err := mr.db.Where("user_id = ? AND DATE_TRUNC('month', updated_at) = ?", userId, startOfMonth).
		Find(moneyManagement).Error; err != nil {
		return err
	}

	return nil
}

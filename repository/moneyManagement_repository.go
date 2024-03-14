package repository

import (
	"merchandise-review-list-backend/model"
	"time"

	"gorm.io/gorm"
)

type IMoneyManagementRepository interface {
	CreateMoneyManagement(moneyManagement *model.MoneyManagement) error
	GetMyMoneyManagements(moneyManagement *[]model.MoneyManagement, userId uint, yearMonth time.Time, yearFlag bool) error
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

func (mr *moneyManagementRepository) GetMyMoneyManagements(moneyManagement *[]model.MoneyManagement, userId uint, yearMonth time.Time, yearFlag bool) error {
	startOfMonth := time.Date(yearMonth.Year(), yearMonth.Month(), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0)

	if yearFlag {
		startOfYear := time.Date(yearMonth.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
		endOfYear := startOfYear.AddDate(1, 0, 0)
		if err := mr.db.Where("user_id = ? AND updated_at >= ? AND updated_at < ?", userId, startOfYear, endOfYear).
			Find(moneyManagement).Error; err != nil {
			return err
		}
	} else {
		if err := mr.db.Where("user_id = ? AND updated_at >= ? AND updated_at < ?", userId, startOfMonth, endOfMonth).
			Find(moneyManagement).Error; err != nil {
			return err
		}
	}

	return nil
}

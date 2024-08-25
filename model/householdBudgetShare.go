package model

import "time"

// 家計簿共有テーブル
type HouseholdBudgetShare struct {
	ID                uint            `json:"id" gorm:"primaryKey"`
	User              User            `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId            uint            `json:"user_id" gorm:"not null"`
	HouseholdBudget   HouseholdBudget `json:"household_budget" gorm:"foreignKey:HouseholdBudgetId; constraint:OnDelete:CASCADE"`
	HouseholdBudgetId uint            `json:"household_budget_id" gorm:"not null"`
	CreatedAt         time.Time       `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time       `gorm:"autoUpdateTime"`
}

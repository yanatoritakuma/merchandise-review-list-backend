package model

import "time"

// 家計簿権限テーブル
type HouseholdBudgetAuthority struct {
	ID                uint            `json:"id" gorm:"primaryKey"`
	Authority         uint            `json:"authority" gorm:"not null"`
	User              User            `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId            uint            `json:"user_id" gorm:"not null"`
	HouseholdBudget   HouseholdBudget `json:"household_budget" gorm:"foreignKey:HouseholdBudgetId; constraint:OnDelete:CASCADE"`
	HouseholdBudgetId uint            `json:"household_budget_id" gorm:"not null"`
	CreatedAt         time.Time       `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time       `gorm:"autoUpdateTime"`
}

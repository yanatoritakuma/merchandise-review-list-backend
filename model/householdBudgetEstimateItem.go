package model

import "time"

// 家計簿予算アイテムテーブル
type HouseholdBudgetEstimateItem struct {
	ID                uint                    `json:"id" gorm:"primaryKey"`
	User              User                    `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId            uint                    `json:"user_id" gorm:"not null"`
	Year              uint                    `json:"year" gorm:"not null"`
	Month             uint                    `json:"month" gorm:"not null"`
	Name              string                  `json:"name" gorm:"not null"`
	Category          HouseholdBudgetCategory `json:"household_budget_category" gorm:"foreignKey:CategoryId; constraint:OnDelete:CASCADE"`
	CategoryId        uint                    `json:"category_id" gorm:"not null"`
	Amount            uint                    `json:"amount" gorm:"not null"`
	FixedCost         bool                    `json:"fixed_cost" gorm:"not null"`
	HouseholdBudgetId uint                    `json:"household_budget_id" gorm:"not null"`
	CreatedAt         time.Time               `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time               `gorm:"autoUpdateTime"`
}

type HouseholdBudgetEstimateItemResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

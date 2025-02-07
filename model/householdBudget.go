package model

import "time"

type HouseholdBudget struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type HouseholdBudgetResponse struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

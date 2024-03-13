package model

import "time"

type MoneyManagement struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title" gorm:"not null"`
	Category   string    `json:"category" gorm:"not null"`
	UnitPrice  uint      `json:"unit_price" gorm:"not null"`
	Quantity   uint      `json:"quantity" gorm:"not null"`
	TotalPrice uint      `json:"total_price" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	User       User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId     uint      `json:"user_id" gorm:"not null"`
}

type MoneyManagementResponse struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Category   string    `json:"category"`
	UnitPrice  uint      `json:"unit_price"`
	Quantity   uint      `json:"quantity"`
	TotalPrice uint      `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
}

type MoneyManagementByCategoryResponse struct {
	Food          []MoneyManagementResponse `json:"food"`
	Drink         []MoneyManagementResponse `json:"drink"`
	Book          []MoneyManagementResponse `json:"book"`
	Fashion       []MoneyManagementResponse `json:"fashion"`
	Furniture     []MoneyManagementResponse `json:"furniture"`
	GamesToys     []MoneyManagementResponse `json:"gamesToys"`
	Beauty        []MoneyManagementResponse `json:"beauty"`
	EveryDayItems []MoneyManagementResponse `json:"everyDayItems"`
	Other         []MoneyManagementResponse `json:"other"`
}

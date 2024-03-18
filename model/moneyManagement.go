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
	UpdatedAt  time.Time `json:"updated_at"`
}

type MoneyManagementByCategoryItemResponse struct {
	Items          []MoneyManagementResponse `json:"items"`
	ItemTotalPrice uint                      `json:"itemTotalPrice"`
}

type MoneyManagementByCategoryResponse struct {
	Food          MoneyManagementByCategoryItemResponse `json:"food"`
	Drink         MoneyManagementByCategoryItemResponse `json:"drink"`
	Book          MoneyManagementByCategoryItemResponse `json:"book"`
	Fashion       MoneyManagementByCategoryItemResponse `json:"fashion"`
	Furniture     MoneyManagementByCategoryItemResponse `json:"furniture"`
	GamesToys     MoneyManagementByCategoryItemResponse `json:"gamesToys"`
	Beauty        MoneyManagementByCategoryItemResponse `json:"beauty"`
	EveryDayItems MoneyManagementByCategoryItemResponse `json:"everyDayItems"`
	Other         MoneyManagementByCategoryItemResponse `json:"other"`
	TotalPrice    uint                                  `json:"totalPrice"`
}

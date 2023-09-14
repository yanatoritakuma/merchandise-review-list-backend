package model

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Stock       bool      `json:"stock" gorm:"not null"`
	Price       uint      `json:"price" gorm:"not null"`
	Review      float64   `json:"review" gorm:"not null"`
	Url         string    `json:"url" gorm:"not null"`
	Image       string    `json:"image" gorm:"not null"`
	Code        string    `json:"code" gorm:"not null"`
	User        User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId      uint      `json:"user_id" gorm:"not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type ProductResponse struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Stock       bool      `json:"stock" gorm:"not null"`
	Price       uint      `json:"price" gorm:"not null"`
	Review      float64   `json:"review" gorm:"not null"`
	Url         string    `json:"url" gorm:"not null"`
	Image       string    `json:"image" gorm:"not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

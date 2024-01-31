package model

import "time"

type ReviewPost struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Text      string    `json:"text" gorm:"not null"`
	Image     string    `json:"image"`
	Review    float64   `json:"review" gorm:"not null"`
	Category  string    `json:"category" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

type ReviewPostResponse struct {
	ID             uint                   `json:"id" gorm:"primaryKey"`
	Title          string                 `json:"title" gorm:"not null"`
	Text           string                 `json:"text" gorm:"not null"`
	Image          string                 `json:"image"`
	Review         float64                `json:"review" gorm:"not null"`
	Category       string                 `json:"category" gorm:"not null"`
	CreatedAt      time.Time              `json:"created_at"`
	User           ReviewPostUserResponse `json:"reviewPostUserResponse"`
	UserId         uint                   `json:"user_id" gorm:"not null"`
	LikeCount      uint                   `json:"like_count"`
	LikeId         uint                   `json:"like_id"`
	LikePostUserId uint                   `json:"like_post_user_id" gorm:"not null"`
	CommentCount   uint                   `json:"comment_count"`
}

type ReviewPostUserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

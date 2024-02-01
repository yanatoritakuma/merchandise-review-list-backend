package model

import "time"

type Comment struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Text       string     `json:"text" gorm:"not null"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	ReviewPost ReviewPost `json:"reviewPost" gorm:"foreignKey:PostId; constraint:OnDelete:CASCADE"`
	PostId     uint       `json:"post_id" gorm:"not null"`
	User       User       `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId     uint       `json:"user_id" gorm:"not null"`
}

type CommentResponse struct {
	ID        uint        `json:"id"`
	Text      string      `json:"text" gorm:"not null"`
	User      CommentUser `json:"comment_user"`
	UserId    uint        `json:"user_id" gorm:"not null"`
	CreatedAt time.Time   `json:"created_at"`
}

type CommentUser struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

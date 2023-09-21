package repository

import (
	"merchandise-review-list-backend/model"

	"gorm.io/gorm"
)

type IReviewPostRepository interface {
	CreateReviewPost(post *model.ReviewPost) error
}

type reviewPostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) IReviewPostRepository {
	return &reviewPostRepository{db}
}

func (pr *reviewPostRepository) CreateReviewPost(reviewPost *model.ReviewPost) error {
	if err := pr.db.Create(reviewPost).Error; err != nil {
		return err
	}
	return nil
}

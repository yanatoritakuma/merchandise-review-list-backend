package repository

import (
	"merchandise-review-list-backend/model"

	"gorm.io/gorm"
)

type IReviewPostRepository interface {
	CreateReviewPost(reviewPost *model.ReviewPost) error
	GetMyReviewPosts(reviewPost *[]model.ReviewPost, userId uint, page int, pageSize int) (int, error)
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

func (pr *reviewPostRepository) GetReviewPostsByIds(reviewPost *[]model.ReviewPost, postIds []uint) error {
	if err := pr.db.Where("id IN (?)", postIds).Order("created_at DESC").Find(reviewPost).Error; err != nil {
		return err
	}
	return nil
}

func (pr *reviewPostRepository) GetMyReviewPosts(reviewPost *[]model.ReviewPost, userId uint, page int, pageSize int) (int, error) {
	offset := (page - 1) * pageSize
	var totalCount int64

	if err := pr.db.Model(&model.ReviewPost{}).Where("user_id=?", userId).Count(&totalCount).Error; err != nil {
		return 0, err
	}

	if err := pr.db.Joins("User").Where("user_id=?", userId).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(reviewPost).Error; err != nil {
		return 0, err
	}
	return int(totalCount), nil
}

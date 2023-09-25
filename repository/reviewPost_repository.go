package repository

import (
	"fmt"
	"merchandise-review-list-backend/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IReviewPostRepository interface {
	CreateReviewPost(reviewPost *model.ReviewPost) error
	UpdateReviewPost(reviewPost *model.ReviewPost, userId uint, postId uint) error

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

func (pr *reviewPostRepository) UpdateReviewPost(reviewPost *model.ReviewPost, userId uint, postId uint) error {
	result := pr.db.Model(reviewPost).Clauses(clause.Returning{}).Where("id=? AND user_id=?", postId, userId).Updates(map[string]interface{}{
		"title":  reviewPost.Title,
		"text":   reviewPost.Text,
		"image":  reviewPost.Image,
		"review": reviewPost.Review,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
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

package repository

import (
	"merchandise-review-list-backend/model"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	CreateComment(comment *model.Comment) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &commentRepository{db}
}

func (cr *commentRepository) CreateComment(comment *model.Comment) error {
	if err := cr.db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

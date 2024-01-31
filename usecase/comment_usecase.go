package usecase

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/repository"
	"merchandise-review-list-backend/validator"
)

type ICommentUsecase interface {
	CreateComment(comment model.Comment) (model.CommentResponse, error)
}

type commentUsecase struct {
	cr repository.ICommentRepository
	cv validator.ICommentValidator
}

func NewCommentUsecase(cr repository.ICommentRepository, cv validator.ICommentValidator) ICommentUsecase {
	return &commentUsecase{cr, cv}
}

func (cu *commentUsecase) CreateComment(comment model.Comment) (model.CommentResponse, error) {
	if err := cu.cv.CommentValidator(comment); err != nil {
		return model.CommentResponse{}, err
	}

	if err := cu.cr.CreateComment(&comment); err != nil {
		return model.CommentResponse{}, err
	}

	resComment := model.CommentResponse{
		ID:     comment.ID,
		UserId: comment.UserId,
	}
	return resComment, nil
}

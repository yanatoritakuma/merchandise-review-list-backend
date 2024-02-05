package usecase

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/repository"
	"merchandise-review-list-backend/validator"
)

type ICommentUsecase interface {
	CreateComment(comment model.Comment) (model.CommentResponse, error)
	DeleteComment(userId uint, id uint) error
	GetCommentsByPostId(postId uint, page int, pageSize int) ([]model.CommentResponse, int, error)
}

type commentUsecase struct {
	cr repository.ICommentRepository
	cv validator.ICommentValidator
	rr repository.IReviewPostRepository
}

func NewCommentUsecase(cr repository.ICommentRepository, cv validator.ICommentValidator, rr repository.IReviewPostRepository) ICommentUsecase {
	return &commentUsecase{cr, cv, rr}
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

func (cu *commentUsecase) DeleteComment(userId uint, id uint) error {
	if err := cu.cr.DeleteComment(userId, id); err != nil {
		return err
	}
	return nil
}

func (cu *commentUsecase) GetCommentsByPostId(postId uint, page int, pageSize int) ([]model.CommentResponse, int, error) {
	comments := []model.Comment{}

	totalCount, err := cu.cr.GetCommentsByPostId(&comments, postId, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	resCounts := []model.CommentResponse{}

	for _, v := range comments {
		user, err := cu.rr.GetUserById(v.UserId)
		if err != nil {
			return nil, 0, err
		}

		c := model.CommentResponse{
			ID:   v.ID,
			Text: v.Text,
			User: model.CommentUser{
				ID:    user.ID,
				Name:  user.Name,
				Image: user.Image,
			},
			UserId:    v.UserId,
			CreatedAt: v.CreatedAt,
		}
		resCounts = append(resCounts, c)
	}
	return resCounts, totalCount, nil
}

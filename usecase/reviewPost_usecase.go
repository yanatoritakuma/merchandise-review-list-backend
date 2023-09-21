package usecase

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/repository"
	"merchandise-review-list-backend/validator"
)

type IReviewPostUsecase interface {
	CreateReviewPost(reviewPost model.ReviewPost) (model.ReviewPostResponse, error)
}

type reviewPostUsecase struct {
	rr repository.IReviewPostRepository
	rv validator.IReviewPostValidator
}

func NewReviewPostUsecase(
	rr repository.IReviewPostRepository,
	rv validator.IReviewPostValidator,
) IReviewPostUsecase {
	return &reviewPostUsecase{rr, rv}
}

func (ru *reviewPostUsecase) CreateReviewPost(reviewPost model.ReviewPost) (model.ReviewPostResponse, error) {
	if err := ru.rv.ReviewPostValidator(reviewPost); err != nil {
		return model.ReviewPostResponse{}, err
	}
	if err := ru.rr.CreateReviewPost(&reviewPost); err != nil {
		return model.ReviewPostResponse{}, err
	}
	resReviewPost := model.ReviewPostResponse{
		ID:        reviewPost.ID,
		Title:     reviewPost.Title,
		Text:      reviewPost.Text,
		Image:     reviewPost.Image,
		CreatedAt: reviewPost.CreatedAt,
		User: model.ReviewPostUserResponse{
			ID:    reviewPost.User.ID,
			Name:  reviewPost.User.Name,
			Image: reviewPost.User.Image,
		},
		UserId: reviewPost.UserId,
	}
	return resReviewPost, nil
}

package usecase

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/repository"
	"merchandise-review-list-backend/validator"
)

type IReviewPostUsecase interface {
	CreateReviewPost(reviewPost model.ReviewPost) (model.ReviewPostResponse, error)
	GetReviewPostsByIds(userId uint, page int, pageSize int) ([]model.ReviewPostResponse, int, error)
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

func (ru *reviewPostUsecase) GetReviewPostsByIds(userId uint, page int, pageSize int) ([]model.ReviewPostResponse, int, error) {
	reviewPosts := []model.ReviewPost{}
	totalCount, err := ru.rr.GetMyReviewPosts(&reviewPosts, userId, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	resReviewPosts := []model.ReviewPostResponse{}
	for _, v := range reviewPosts {

		r := model.ReviewPostResponse{
			ID:        v.ID,
			Title:     v.Title,
			Text:      v.Text,
			Image:     v.Image,
			Review:    v.Review,
			CreatedAt: v.CreatedAt,
			User: model.ReviewPostUserResponse{
				ID:    v.User.ID,
				Name:  v.User.Name,
				Image: v.User.Image,
			},
			UserId: v.UserId,
		}
		resReviewPosts = append(resReviewPosts, r)
	}
	return resReviewPosts, totalCount, nil
}

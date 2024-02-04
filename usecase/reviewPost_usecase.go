package usecase

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/repository"
	"merchandise-review-list-backend/validator"
)

type IReviewPostUsecase interface {
	CreateReviewPost(reviewPost model.ReviewPost) (model.ReviewPostResponse, error)
	UpdateReviewPost(reviewPost model.ReviewPost, userId uint, postId uint) (model.ReviewPostResponse, error)
	DeleteReviewPost(userId uint, postId uint) error
	GetMyReviewPosts(userId uint, page int, pageSize int) ([]model.ReviewPostResponse, int, error)
	GetReviewPostById(postId uint) (model.ReviewPostResponse, error)
	GetReviewPostLists(category string, page int, pageSize int, userId uint) ([]model.ReviewPostResponse, int, error)
	GetMyLikes(userId uint, page int, pageSize int) ([]model.ReviewPostResponse, int, error)
}

type reviewPostUsecase struct {
	rr repository.IReviewPostRepository
	rv validator.IReviewPostValidator
	lr repository.ILikeRepository
}

func NewReviewPostUsecase(
	rr repository.IReviewPostRepository,
	rv validator.IReviewPostValidator,
	lr repository.ILikeRepository,
) IReviewPostUsecase {
	return &reviewPostUsecase{rr, rv, lr}
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
		Category:  reviewPost.Category,
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

func (ru *reviewPostUsecase) UpdateReviewPost(reviewPost model.ReviewPost, userId uint, postId uint) (model.ReviewPostResponse, error) {
	if err := ru.rv.ReviewPostValidator(reviewPost); err != nil {
		return model.ReviewPostResponse{}, err
	}
	if err := ru.rr.UpdateReviewPost(&reviewPost, userId, postId); err != nil {
		return model.ReviewPostResponse{}, err
	}
	resReviewPost := model.ReviewPostResponse{
		ID:        reviewPost.ID,
		Title:     reviewPost.Title,
		Text:      reviewPost.Text,
		Image:     reviewPost.Image,
		Review:    reviewPost.Review,
		Category:  reviewPost.Category,
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

func (ru *reviewPostUsecase) DeleteReviewPost(userId uint, postId uint) error {
	if err := ru.rr.DeleteReviewPost(userId, postId); err != nil {
		return err
	}
	return nil
}

func (ru *reviewPostUsecase) GetMyReviewPosts(userId uint, page int, pageSize int) ([]model.ReviewPostResponse, int, error) {
	reviewPosts := []model.ReviewPost{}
	totalCount, err := ru.rr.GetMyReviewPosts(&reviewPosts, userId, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	resReviewPosts := []model.ReviewPostResponse{}
	for _, v := range reviewPosts {

		likes := []model.Like{}
		err = ru.rr.GetLikesByPostId(&likes, v.ID)
		if err != nil {
			return nil, 0, err
		}

		likeCount := uint(len(likes))
		likeId := uint(0)

		for _, like := range likes {
			if like.UserId == userId {
				likeId = uint(like.ID)
			}
		}

		r := model.ReviewPostResponse{
			ID:        v.ID,
			Title:     v.Title,
			Text:      v.Text,
			Image:     v.Image,
			Review:    v.Review,
			Category:  v.Category,
			CreatedAt: v.CreatedAt,
			User: model.ReviewPostUserResponse{
				ID:    v.User.ID,
				Name:  v.User.Name,
				Image: v.User.Image,
			},
			UserId:    v.UserId,
			LikeCount: likeCount,
			LikeId:    likeId,
		}
		resReviewPosts = append(resReviewPosts, r)
	}
	return resReviewPosts, totalCount, nil
}

func (ru *reviewPostUsecase) GetReviewPostById(postId uint) (model.ReviewPostResponse, error) {
	reviewPost := model.ReviewPost{}
	if err := ru.rr.GetReviewPostById(&reviewPost, postId); err != nil {
		return model.ReviewPostResponse{}, err
	}
	user, err := ru.rr.GetUserById(reviewPost.UserId)
	if err != nil {
		return model.ReviewPostResponse{}, err
	}
	resReviewPost := model.ReviewPostResponse{
		ID:        reviewPost.ID,
		Title:     reviewPost.Title,
		Text:      reviewPost.Text,
		Image:     reviewPost.Image,
		Review:    reviewPost.Review,
		Category:  reviewPost.Category,
		CreatedAt: reviewPost.CreatedAt,
		User: model.ReviewPostUserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Image: user.Image,
		},
		UserId: reviewPost.UserId,
	}
	return resReviewPost, nil
}

func (ru *reviewPostUsecase) GetReviewPostLists(category string, page int, pageSize int, userId uint) ([]model.ReviewPostResponse, int, error) {
	reviewPosts := []model.ReviewPost{}

	totalCount, err := ru.rr.GetReviewPostLists(&reviewPosts, category, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	resReviewPosts := []model.ReviewPostResponse{}
	for _, v := range reviewPosts {
		user, err := ru.rr.GetUserById(v.UserId)
		if err != nil {
			return nil, 0, err
		}

		likes := []model.Like{}
		err = ru.rr.GetLikesByPostId(&likes, v.ID)
		if err != nil {
			return nil, 0, err
		}

		likeCount := uint(len(likes))
		likeId := uint(0)
		for _, like := range likes {
			if like.UserId == userId {
				likeId = uint(like.ID)
			}
		}

		comments := []model.Comment{}
		err = ru.rr.GetCommentsByPostId(&comments, v.ID)
		if err != nil {
			return nil, 0, err
		}

		commentCount := uint(len(comments))

		r := model.ReviewPostResponse{
			ID:        v.ID,
			Title:     v.Title,
			Text:      v.Text,
			Image:     v.Image,
			Review:    v.Review,
			Category:  v.Category,
			CreatedAt: v.CreatedAt,
			User: model.ReviewPostUserResponse{
				ID:    user.ID,
				Name:  user.Name,
				Image: user.Image,
			},
			UserId:       v.UserId,
			LikeCount:    likeCount,
			LikeId:       likeId,
			CommentCount: commentCount,
		}

		resReviewPosts = append(resReviewPosts, r)
	}
	return resReviewPosts, totalCount, nil
}

func (ru *reviewPostUsecase) GetMyLikes(userId uint, page int, pageSize int) ([]model.ReviewPostResponse, int, error) {
	totalLikeCount, err := ru.lr.GetMyLikeCount(userId)
	if err != nil {
		return nil, 0, err
	}

	postIds, err := ru.lr.GetMyLikePostIdsByUserId(userId, page, pageSize)
	resLikePosts := []model.ReviewPostResponse{}
	for _, v := range postIds {
		likes := []model.Like{}
		err = ru.rr.GetLikesByPostId(&likes, v)
		if err != nil {
			return nil, 0, err
		}

		likeCount := uint(len(likes))
		likeId := uint(0)
		for _, like := range likes {
			if like.UserId == userId {
				likeId = uint(like.ID)
			}
		}

		post := model.ReviewPost{}
		if err := ru.rr.GetReviewPostById(&post, v); err != nil {
			return nil, 0, err
		}

		user, err := ru.rr.GetUserById(post.UserId)
		if err != nil {
			return nil, 0, err
		}

		p := model.ReviewPostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Text:      post.Text,
			Image:     post.Image,
			Review:    post.Review,
			Category:  post.Category,
			CreatedAt: post.CreatedAt,
			User: model.ReviewPostUserResponse{
				ID:    user.ID,
				Name:  user.Name,
				Image: user.Image,
			},
			UserId:    post.UserId,
			LikeCount: likeCount,
			LikeId:    likeId,
		}
		resLikePosts = append(resLikePosts, p)
	}
	return resLikePosts, totalLikeCount, nil
}

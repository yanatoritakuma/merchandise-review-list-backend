package validator

import (
	"merchandise-review-list-backend/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IReviewPostValidator interface {
	ReviewPostValidator(reviewPost model.ReviewPost) error
}

type reviewPostValidator struct{}

func NewReviewPostValidator() IReviewPostValidator {
	return &reviewPostValidator{}
}

func (rv *reviewPostValidator) ReviewPostValidator(reviewPost model.ReviewPost) error {
	return validation.ValidateStruct(&reviewPost,
		validation.Field(
			&reviewPost.Title,
			validation.Required.Error("text is required"),
			validation.RuneLength(1, 50).Error("limites max 50 char"),
		),
		validation.Field(
			&reviewPost.Text,
			validation.Required.Error("text is required"),
			validation.RuneLength(1, 150).Error("limites max 150 char"),
		),
		validation.Field(
			&reviewPost.Review,
			validation.Required.Error("review is required"),
		),
	)
}

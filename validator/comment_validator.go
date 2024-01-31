package validator

import (
	"merchandise-review-list-backend/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ICommentValidator interface {
	CommentValidator(comment model.Comment) error
}

type commentValidator struct{}

func NewCommentValidator() ICommentValidator {
	return &commentValidator{}
}

func (cv *commentValidator) CommentValidator(comment model.Comment) error {
	return validation.ValidateStruct(&comment,
		validation.Field(
			&comment.Text,
			validation.Required.Error("text is required"),
			validation.RuneLength(1, 150).Error("limites max 150 char"),
		),
	)
}

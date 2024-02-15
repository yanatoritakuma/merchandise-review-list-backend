package validator

import (
	"merchandise-review-list-backend/model"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IProductValidator interface {
	ProductValidator(product model.Product) error
}

type productValidator struct{}

func NewProductValidator() IProductValidator {
	return &productValidator{}
}

func (pv *productValidator) ProductValidator(product model.Product) error {
	return validation.ValidateStruct(&product,
		validation.Field(
			&product.TimeLimit,

			validation.When(product.TimeLimit != time.Time{},
				validation.By(func(value interface{}) error {
					timeLimit, ok := value.(time.Time)
					if !ok {
						return validation.NewError("validation_type", "invalid TimeLimit type")
					}

					now := time.Now()

					// 現在日時よりも過去の場合はエラー
					if timeLimit.Before(now) {
						return validation.NewError("validation_future", "TimeLimit must be in the future")
					}

					return nil
				}),
			),
		),
	)
}

package validator

import (
	"merchandise-review-list-backend/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IMoneyManagementValidator interface {
	MoneyManagementValidator(moneyManagement model.MoneyManagement) error
}

type moneyManagementValidator struct{}

func NewMoneyManagementValidator() IMoneyManagementValidator {
	return &moneyManagementValidator{}
}

func (mv *moneyManagementValidator) MoneyManagementValidator(moneyManagement model.MoneyManagement) error {
	return validation.ValidateStruct(&moneyManagement,
		validation.Field(
			&moneyManagement.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 50).Error("limites max 50 char"),
		),
		validation.Field(
			&moneyManagement.Category,
			validation.Required.Error("category is required"),
		),
		validation.Field(
			&moneyManagement.UnitPrice,
			validation.Required.Error("unitPrice is required"),
		),
		validation.Field(
			&moneyManagement.Quantity,
			validation.Required.Error("quantity is required"),
		),
		validation.Field(
			&moneyManagement.TotalPrice,
			validation.Required.Error("totalPrice is required"),
		),
		validation.Field(
			&moneyManagement.UpdatedAt,
			validation.Required.Error("updatedAt is required"),
		),
	)
}

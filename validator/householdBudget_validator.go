package validator

import (
	"merchandise-review-list-backend/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IHouseholdBudgetValidator interface {
	HouseholdBudgetValidator(householdBudget model.HouseholdBudget) error
}

type householdBudgetValidator struct{}

func NewHouseholdBudgetValidator() IHouseholdBudgetValidator {
	return &householdBudgetValidator{}
}

func (hv *householdBudgetValidator) HouseholdBudgetValidator(householdBudget model.HouseholdBudget) error {
	return validation.ValidateStruct(&householdBudget,
		validation.Field(
			&householdBudget.Title,
			validation.Required.Error("text is required"),
			validation.RuneLength(1, 50).Error("limites max 50 char"),
		),
	)
}

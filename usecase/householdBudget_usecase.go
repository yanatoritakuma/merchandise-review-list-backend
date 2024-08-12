package usecase

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/repository"
	"merchandise-review-list-backend/validator"
)

type IHouseholdBudgetUsecase interface {
	CreateHouseholdBudget(householdBudget model.HouseholdBudget) (model.HouseholdBudgetUsecaseResponse, error)
}

type householdBudgetUsecase struct {
	hr repository.IHouseholdBudgetRepository
	hv validator.IHouseholdBudgetValidator
}

func NweHouseholdBudgetUsecase(hr repository.IHouseholdBudgetRepository, hv validator.IHouseholdBudgetValidator) IHouseholdBudgetUsecase {
	return &householdBudgetUsecase{hr, hv}
}

// 家計簿を作成する
func (hu *householdBudgetUsecase) CreateHouseholdBudget(householdBudget model.HouseholdBudget) (model.HouseholdBudgetUsecaseResponse, error) {
	if err := hu.hv.HouseholdBudgetValidator(householdBudget); err != nil {
		return model.HouseholdBudgetUsecaseResponse{}, err
	}

	if err := hu.hr.CreateHouseholdBudget(&householdBudget); err != nil {
		return model.HouseholdBudgetUsecaseResponse{}, err
	}

	resHouseholdBudget := model.HouseholdBudgetUsecaseResponse{
		ID:     householdBudget.ID,
		UserId: householdBudget.UserId,
		Title:  householdBudget.Title,
	}

	return resHouseholdBudget, nil
}

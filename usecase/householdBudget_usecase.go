package usecase

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/repository"
	"merchandise-review-list-backend/validator"
)

type IHouseholdBudgetUsecase interface {
	CreateHouseholdBudget(householdBudget model.HouseholdBudget) (model.HouseholdBudgetResponse, error)
	GetMyHouseholdBudget(userId uint, page int, pageSize int) ([]model.HouseholdBudgetResponse, int, error)
}

type householdBudgetUsecase struct {
	hr repository.IHouseholdBudgetRepository
	hv validator.IHouseholdBudgetValidator
}

func NweHouseholdBudgetUsecase(hr repository.IHouseholdBudgetRepository, hv validator.IHouseholdBudgetValidator) IHouseholdBudgetUsecase {
	return &householdBudgetUsecase{hr, hv}
}

// 家計簿を作成する
func (hu *householdBudgetUsecase) CreateHouseholdBudget(householdBudget model.HouseholdBudget) (model.HouseholdBudgetResponse, error) {
	if err := hu.hv.HouseholdBudgetValidator(householdBudget); err != nil {
		return model.HouseholdBudgetResponse{}, err
	}

	if err := hu.hr.CreateHouseholdBudget(&householdBudget); err != nil {
		return model.HouseholdBudgetResponse{}, err
	}

	resHouseholdBudget := model.HouseholdBudgetResponse{
		ID:     householdBudget.ID,
		UserId: householdBudget.UserId,
		Title:  householdBudget.Title,
	}

	return resHouseholdBudget, nil
}

// 自分が作成した家計簿取得 todo: 家計簿共有テーブル（household_budget_share）を参照して自分が所属している家計も取得
func (hu *householdBudgetUsecase) GetMyHouseholdBudget(userId uint, page int, pageSize int) ([]model.HouseholdBudgetResponse, int, error) {
	householdBudgets := []model.HouseholdBudget{}

	totalCount, err := hu.hr.GetMyHouseholdBudget(&householdBudgets, userId, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	resHouseholdBudgets := []model.HouseholdBudgetResponse{}

	for _, v := range householdBudgets {
		h := model.HouseholdBudgetResponse{
			ID:        v.ID,
			UserId:    v.UserId,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
		}

		resHouseholdBudgets = append(resHouseholdBudgets, h)
	}

	return resHouseholdBudgets, totalCount, nil
}

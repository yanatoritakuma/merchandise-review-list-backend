package controller

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IHouseholdBudgetEstimateItemController interface {
	CreateHouseholdBudgetEstimateItem(c echo.Context) error
}

type householdBudgetEstimateItemController struct {
	hu usecase.IHouseholdBudgetEstimateItemUsecase
}

func NewHouseholdBudgetEstimateItemController(hu usecase.IHouseholdBudgetEstimateItemUsecase) IHouseholdBudgetEstimateItemController {
	return &householdBudgetEstimateItemController{hu}
}

func (hc *householdBudgetEstimateItemController) CreateHouseholdBudgetEstimateItem(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	householdBudgetEstimateItem := model.HouseholdBudgetEstimateItem{}
	if err := c.Bind(&householdBudgetEstimateItem); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	householdBudgetEstimateItem.UserId = uint(userId.(float64))

	householdBudgetEstimateItemRes, err := hc.hu.CreateHouseholdBudgetEstimateItem(householdBudgetEstimateItem)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, householdBudgetEstimateItemRes)
}

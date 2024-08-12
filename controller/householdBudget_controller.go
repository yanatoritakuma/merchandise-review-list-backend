package controller

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IHouseholdBudgetController interface {
	CreateHouseholdBudget(c echo.Context) error
	GetMyHouseholdBudget(c echo.Context) error
}

type householdBudgetController struct {
	hu usecase.IHouseholdBudgetUsecase
}

func NewHouseholdBudgetController(hu usecase.IHouseholdBudgetUsecase) IHouseholdBudgetController {
	return &householdBudgetController{hu}
}

func (hc *householdBudgetController) CreateHouseholdBudget(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	householdBudget := model.HouseholdBudget{}
	if err := c.Bind(&householdBudget); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	householdBudget.UserId = uint(userId.(float64))

	householdBudgetRes, err := hc.hu.CreateHouseholdBudget(householdBudget)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, householdBudgetRes)
}

func (hc *householdBudgetController) GetMyHouseholdBudget(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	householdBudgetRes, totalPageCount, err := hc.hu.GetMyHouseholdBudget(uint(userId.(float64)), page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"totalPageCount":   totalPageCount,
		"householdBudgets": householdBudgetRes,
	}

	return c.JSON(http.StatusOK, response)
}

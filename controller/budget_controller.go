package controller

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IBudgetController interface {
	CreateBudget(c echo.Context) error
	UpdateBudget(c echo.Context) error
	GetBudgetByUserId(c echo.Context) error
}

type budgetController struct {
	bu usecase.IBudgetUsecase
}

func NewBudgetController(bu usecase.IBudgetUsecase) IBudgetController {
	return &budgetController{bu}
}

func (bc *budgetController) CreateBudget(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	budget := model.Budget{}
	if err := c.Bind(&budget); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	budget.UserId = uint(userId.(float64))

	budgetRes, err := bc.bu.CreateProduct(budget)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, budgetRes)
}

func (bc *budgetController) UpdateBudget(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("id")
	budgetId, _ := strconv.Atoi(id)

	budget := model.Budget{}

	if err := c.Bind(&budget); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	budgetRes, err := bc.bu.UpdateBudget(budget, uint(userId.(float64)), uint(budgetId))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, budgetRes)
}

func (bc *budgetController) GetBudgetByUserId(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	year := c.QueryParam("year")
	month := c.QueryParam("month")

	budgetRes, err := bc.bu.GetBudgetByUserId(uint(userId.(float64)), year, month)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"budget": budgetRes,
	}

	return c.JSON(http.StatusOK, response)
}

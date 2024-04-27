package controller

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IMoneyManagementController interface {
	CreateMoneyManagement(c echo.Context) error
	UpdateMoneyManagement(c echo.Context) error
	DeleteMoneyManagement(c echo.Context) error
	GetMyMoneyManagements(c echo.Context) error
}

type moneyManagementController struct {
	mu usecase.IMoneyManagementUsecase
}

func NewMoneyManagementController(mu usecase.IMoneyManagementUsecase) IMoneyManagementController {
	return &moneyManagementController{mu}
}

func (mc *moneyManagementController) CreateMoneyManagement(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	moneyManagement := model.MoneyManagement{}
	if err := c.Bind(&moneyManagement); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	moneyManagement.UserId = uint(userId.(float64))
	moneyManagementRes, err := mc.mu.CreateMoneyManagement(moneyManagement)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, moneyManagementRes)
}

func (mc *moneyManagementController) UpdateMoneyManagement(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("id")
	managementId, _ := strconv.Atoi(id)

	moneyManagement := model.MoneyManagement{}

	if err := c.Bind(&moneyManagement); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	moneyManagementRes, err := mc.mu.UpdateMoneyManagement(moneyManagement, uint(userId.(float64)), uint(managementId))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, moneyManagementRes)
}

func (mc *moneyManagementController) DeleteMoneyManagement(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("id")
	managementId, _ := strconv.Atoi(id)

	err := mc.mu.DeleteMoneyManagement(uint(userId.(float64)), uint(managementId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (mc *moneyManagementController) GetMyMoneyManagements(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	yearMonth, err := strconv.Atoi(c.QueryParam("yearMonth"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid yearMonth format")
	}
	yearFlag, err := strconv.ParseBool(c.QueryParam("yearFlag"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid yearFlag format")
	}

	// yearMonthをtime.Timeに変換する
	year := yearMonth / 100  // 年を取得
	month := yearMonth % 100 // 月を取得

	yearMonthTime := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	moneyManagementsRes, err := mc.mu.GetMyMoneyManagements(uint(userId.(float64)), yearMonthTime, yearFlag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, moneyManagementsRes)
}

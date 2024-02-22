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

type IProductController interface {
	CreateProduct(c echo.Context) error
	UpdateTimeLimit(c echo.Context) error
	DeleteProduct(c echo.Context) error
	GetMyProducts(c echo.Context) error
	GetMyProductsTimeLimitAll(c echo.Context) error
	GetMyProductsTimeLimitYearMonth(c echo.Context) error
}

type productController struct {
	pu usecase.IProductUsecase
}

func NewProductController(pu usecase.IProductUsecase) IProductController {
	return &productController{pu}
}

func (pc *productController) CreateProduct(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	product := model.Product{}
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	product.UserId = uint(userId.(float64))
	productRes, err := pc.pu.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, productRes)
}

func (pc *productController) UpdateTimeLimit(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("productId")
	productId, _ := strconv.Atoi(id)

	product := model.Product{}
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	productRes, err := pc.pu.UpdateTimeLimit(product, uint(userId.(float64)), uint(productId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, productRes)
}

func (pc *productController) DeleteProduct(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("productId")
	productId, _ := strconv.Atoi(id)

	err := pc.pu.DeleteProduct(uint(userId.(float64)), uint(productId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (pc *productController) GetMyProducts(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	productsRes, totalPageCount, err := pc.pu.GetMyProducts(uint(userId.(float64)), page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"totalPageCount": totalPageCount,
		"products":       productsRes,
	}

	return c.JSON(http.StatusOK, response)

}

func (pc *productController) GetMyProductsTimeLimitAll(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	productsTimeLimitRes, totalPageCount, err := pc.pu.GetMyProductsTimeLimitAll(uint(userId.(float64)), page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"totalPageCount": totalPageCount,
		"products":       productsTimeLimitRes,
	}

	return c.JSON(http.StatusOK, response)
}

func (pc *productController) GetMyProductsTimeLimitYearMonth(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	yearMonth, err := strconv.Atoi(c.QueryParam("yearMonth"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid yearMonth format")
	}

	// yearMonthをtime.Timeに変換する
	year := yearMonth / 100  // 年を取得
	month := yearMonth % 100 // 月を取得

	yearMonthTime := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	productsTimeLimitRes, err := pc.pu.GetMyProductsTimeLimitYearMonth(uint(userId.(float64)), yearMonthTime)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"productNumbers": productsTimeLimitRes,
	}

	return c.JSON(http.StatusOK, response)
}

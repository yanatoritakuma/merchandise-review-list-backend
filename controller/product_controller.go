package controller

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IProductController interface {
	CreateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
	GetMyProducts(c echo.Context) error
	GetMyProductsTimeLimit(c echo.Context) error
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

func (pc *productController) GetMyProductsTimeLimit(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	productsTimeLimitRes, totalPageCount, err := pc.pu.GetMyProductsTimeLimit(uint(userId.(float64)), page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"totalPageCount":    totalPageCount,
		"productsTimeLimit": productsTimeLimitRes,
	}

	return c.JSON(http.StatusOK, response)
}

package controller

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IReviewPostController interface {
	CreateReviewPost(c echo.Context) error
	GetReviewPostsByIds(c echo.Context) error
}

type reviewPostController struct {
	ru usecase.IReviewPostUsecase
}

func NewReviewPostController(ru usecase.IReviewPostUsecase) IReviewPostController {
	return &reviewPostController{ru}
}

func (rc *reviewPostController) CreateReviewPost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	reviewPost := model.ReviewPost{}
	if err := c.Bind(&reviewPost); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	reviewPost.UserId = uint(userId.(float64))
	reviewPostRes, err := rc.ru.CreateReviewPost(reviewPost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, reviewPostRes)
}

func (rc *reviewPostController) GetReviewPostsByIds(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	reviewPostsRes, totalPageCount, err := rc.ru.GetReviewPostsByIds(uint(userId.(float64)), page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"totalPageCount": totalPageCount,
		"reviewPosts":    reviewPostsRes,
	}

	return c.JSON(http.StatusOK, response)
}

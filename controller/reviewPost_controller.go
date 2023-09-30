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
	UpdateReviewPost(c echo.Context) error
	DeleteReviewPost(c echo.Context) error
	GetMyReviewPosts(c echo.Context) error
	GetReviewPostById(c echo.Context) error
	GetReviewPostLists(c echo.Context) error
	GetMyLikes(c echo.Context) error
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

func (rc *reviewPostController) UpdateReviewPost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("postId")
	postId, _ := strconv.Atoi(id)

	reviewPost := model.ReviewPost{}
	if err := c.Bind(&reviewPost); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	reviewPostRes, err := rc.ru.UpdateReviewPost(reviewPost, uint(userId.(float64)), uint(postId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, reviewPostRes)
}

func (rc *reviewPostController) GetMyReviewPosts(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	reviewPostsRes, totalPageCount, err := rc.ru.GetMyReviewPosts(uint(userId.(float64)), page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"totalPageCount": totalPageCount,
		"reviewPosts":    reviewPostsRes,
	}

	return c.JSON(http.StatusOK, response)
}

func (rc *reviewPostController) GetReviewPostById(c echo.Context) error {
	id := c.Param("postId")
	postId, _ := strconv.Atoi(id)
	reviewPostRes, err := rc.ru.GetReviewPostById(uint(postId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, reviewPostRes)
}

func (rc *reviewPostController) DeleteReviewPost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("postId")
	postId, _ := strconv.Atoi(id)

	err := rc.ru.DeleteReviewPost(uint(userId.(float64)), uint(postId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (rc *reviewPostController) GetReviewPostLists(c echo.Context) error {
	category := c.Param("category")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
	userId, _ := strconv.Atoi(c.QueryParam("userId"))
	reviewPostsRes, totalPageCount, err := rc.ru.GetReviewPostLists(category, page, pageSize, uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"totalPageCount": totalPageCount,
		"reviewPosts":    reviewPostsRes,
	}

	return c.JSON(http.StatusOK, response)
}

func (rc *reviewPostController) GetMyLikes(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	likePostsRes, totalLikeCount, err := rc.ru.GetMyLikes(uint(userId.(float64)), page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"totalPageCount": totalLikeCount,
		"reviewPosts":    likePostsRes,
	}

	return c.JSON(http.StatusOK, response)
}

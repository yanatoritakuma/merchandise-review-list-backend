package controller

import (
	"merchandise-review-list-backend/model"
	"merchandise-review-list-backend/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ICommentController interface {
	CreateComment(c echo.Context) error
	DeleteComment(c echo.Context) error
	GetCommentsByPostId(c echo.Context) error
}

type commentController struct {
	cu usecase.ICommentUsecase
}

func NewCommentController(cu usecase.ICommentUsecase) ICommentController {
	return &commentController{cu}
}

func (cc *commentController) CreateComment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	comment := model.Comment{}
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	comment.User.ID = uint(userId.(float64))
	commentRes, err := cc.cu.CreateComment(comment)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, commentRes)
}

func (cc *commentController) DeleteComment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("id")
	commentId, _ := strconv.Atoi(id)

	err := cc.cu.DeleteComment(uint(userId.(float64)), uint(commentId))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (cc *commentController) GetCommentsByPostId(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
	postId, _ := strconv.Atoi(c.QueryParam("postId"))

	commentsRes, totalPageCount, err := cc.cu.GetCommentsByPostId(uint(postId), page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"totalPageCount": totalPageCount,
		"commentsRes":    commentsRes,
	}

	return c.JSON(http.StatusOK, response)
}

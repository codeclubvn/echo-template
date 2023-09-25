package controller

import (
	"net/http"
	"trail_backend/dto"
	"trail_backend/utils"

	"github.com/labstack/echo/v4"
	"trail_backend/service"
)

type PostController struct {
	BaseController
	postService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}

func (h *PostController) Create(c echo.Context) error {
	var req dto.CreatePostRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	post, err := h.postService.Create(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusCreated, "Success", post)
}

func (h *PostController) Update(c echo.Context) error {
	var req dto.UpdatePostRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	userId := utils.GetUserStringIDFromContext(c)
	req.UserId = userId

	_, err := h.postService.Update(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)

	}

	return h.Response(c, http.StatusOK, "Success", nil)
}

func (h *PostController) GetList(c echo.Context) error {
	var req dto.GetListPostRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	userId := utils.GetUserStringIDFromContext(c)
	req.UserId = userId

	posts, err := h.postService.GetList(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", posts)
}

func (h *PostController) Delete(c echo.Context) error {
	var req dto.DeletePostRequest
	req.Id = utils.ParseStringIDFromUri(c)
	req.UserId = utils.GetUserStringIDFromContext(c)

	err := h.postService.Delete(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", nil)
}

func (h *PostController) GetOne(c echo.Context) error {
	var req dto.GetOnePostRequest
	req.Id = utils.ParseStringIDFromUri(c)

	res, err := h.postService.GetOne(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", res)
}

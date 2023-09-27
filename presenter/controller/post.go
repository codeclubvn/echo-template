package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	utils "trail_backend/pkg/utils"
	"trail_backend/presenter/request"
	"trail_backend/usecase"
)

type PostController struct {
	BaseController
	postService usecase.PostService
}

func NewPostController(postService usecase.PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}

// Create
//
//	@Summary		Create
//	@Description	Create
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Param			Authorization		header		string						true	"authorization token"
//	@Param			CreatePostRequest	body		request.CreatePostRequest	true	"CreatePostRequest"
//	@Success		200					{object}	model.Post					"success"
//	@Router			/v1/api/posts [POST]
func (h *PostController) Create(c echo.Context) error {
	var req request.CreatePostRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	res, err := h.postService.Create(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusCreated, "Success", res)
}

// Update
//
//	@Summary		Update
//	@Description	Update
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Param			Authorization		header		string						true	"authorization token"
//	@Param			UpdatePostRequest	body		request.UpdatePostRequest	true	"UpdatePostRequest"
//	@Success		200					{object}	model.Post					"success"
//	@Router			/v1/api/posts [PUT]
func (h *PostController) Update(c echo.Context) error {
	var req request.UpdatePostRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	res, err := h.postService.Update(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", res)
}

// GetList
//
//	@Summary		GetList
//	@Description	GetList
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Param			Authorization		header		string						true	"authorization token"
//	@Param			GetListPostRequest	body		request.GetListPostRequest	true	"GetListPostRequest"
//	@Success		200					{object}	[]model.Post				"success"
//	@Router			/v1/api/posts [GET]
func (h *PostController) GetList(c echo.Context) error {
	var req request.GetListPostRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	data, total, err := h.postService.GetList(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.ResponseList(c, "Success", total, data)
}

// Delete
//
//	@Summary		Delete
//	@Description	Delete
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string					true	"authorization token"
//	@Param			id				path		string					true	"id"
//	@Success		200				{object}	entity.SimpleResponse	"success"
//	@Router			/v1/api/posts/{id} [DELETE]
func (h *PostController) Delete(c echo.Context) error {
	id := utils.ParseStringIDFromUri(c)
	if err := h.postService.Delete(c.Request().Context(), id); err != nil {
		return h.ResponseError(c, err)
	}
	return h.Response(c, http.StatusOK, "Success", id)
}

// GetOne
//
//	@Summary		GetOne
//	@Description	GetOne
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string		true	"authorization token"
//	@Success		200				{object}	model.Post	"success"
//	@Param			id				path		string		true	"id"
//	@Router			/v1/api/posts/{id} [GET]
func (h *PostController) GetOne(c echo.Context) error {
	id := utils.ParseStringIDFromUri(c)
	res, err := h.postService.GetOne(c.Request().Context(), id)
	if err != nil {
		return h.ResponseError(c, err)
	}
	return h.Response(c, http.StatusOK, "Success", res)
}

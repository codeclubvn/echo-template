package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	utils "trial_backend/pkg/utils"
	"trial_backend/presenter/request"
	"trial_backend/usecase"
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
//	@Security		Authorization
//	@Summary		Create
//	@Description	Create
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Param			CreatePostRequest	body		request.CreatePostRequest	true	"CreatePostRequest"
//	@Success		200					{object}	model.Post					"success"
//	@Router			/v1/api/posts [POST]
func (h *PostController) Create(c echo.Context) error {
	var req request.CreatePostRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidatorError(c, err)
	}

	userId, err := utils.GetUserUUIDFromContext(c)
	if err != nil {
		return h.ResponseValidatorError(c, err)
	}
	req.UserId = userId

	// check files is list uuid
	if err = utils.CheckListUUID(req.Files); err != nil {
		return h.ResponseError(c, err)
	}

	res, err := h.postService.Create(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusCreated, "Success", res)
}

// Update
//
//	@Security		Authorization
//	@Summary		Update
//	@Description	Update
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Param			UpdatePostRequest	body		request.UpdatePostRequest	true	"UpdatePostRequest"
//	@Success		200					{object}	model.Post					"success"
//	@Router			/v1/api/posts [PUT]
func (h *PostController) Update(c echo.Context) error {
	var req request.UpdatePostRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidatorError(c, err)
	}
	userId, err := utils.GetUserUUIDFromContext(c)
	if err != nil {
		return h.ResponseValidatorError(c, err)
	}

	// check files is list uuid
	if err = utils.CheckListUUID(req.Files); err != nil {
		return h.ResponseError(c, err)
	}

	req.UserId = userId
	res, err := h.postService.Update(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", res)
}

// GetList
//
//	@Security		Authorization
//	@Summary		GetList
//	@Description	GetList
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Param			GetListPostRequest	query		request.GetListPostRequest	true	"GetListPostRequest"
//	@Success		200					{object}	[]model.Post				"success"
//	@Router			/v1/api/posts [GET]
func (h *PostController) GetList(c echo.Context) error {
	var req request.GetListPostRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidatorError(c, err)
	}

	data, total, err := h.postService.GetList(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.ResponseList(c, "Success", total, data)
}

// Delete
//
//	@Security		Authorization
//	@Summary		Delete
//	@Description	Delete
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string					true	"id"
//	@Success		200	{object}	entity.SimpleResponse	"success"
//	@Router			/v1/api/posts/{id} [DELETE]
func (h *PostController) Delete(c echo.Context) error {
	id := utils.ParseStringIDFromUri(c)
	userId, err := utils.GetUserUUIDFromContext(c)
	if err != nil {
		return h.ResponseValidatorError(c, err)
	}

	req := request.DeletePostRequest{
		ID:     id,
		UserId: userId,
	}
	if err := h.postService.Delete(c.Request().Context(), req); err != nil {
		return h.ResponseError(c, err)
	}
	return h.Response(c, http.StatusOK, "Success", id)
}

// GetOne
//
//	@Security		Authorization
//	@Summary		GetOne
//	@Description	GetOne
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Post	"success"
//	@Param			id	path		string		true	"id"
//	@Router			/v1/api/posts/{id} [GET]
func (h *PostController) GetOne(c echo.Context) error {
	id := utils.ParseStringIDFromUri(c)
	res, err := h.postService.GetOne(c.Request().Context(), id)
	if err != nil {
		return h.ResponseError(c, err)
	}
	return h.Response(c, http.StatusOK, "Success", res)
}

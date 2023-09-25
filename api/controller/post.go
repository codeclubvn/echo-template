package controller

import (
	"net/http"
	"trail_backend/api/dto"
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

// Create
// @Summary		Create
// @Description	Create
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string								true	"authorization token"
// @Success		200				{object}	dto.SimpleResponse	"success"
// @Router		/v1/api/post [GET]
func (h *PostController) Create(c echo.Context) error {
	var req dto.CreatePostRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}
	userId := utils.GetUserStringIDFromContext(c)
	req.UserId = userId

	data, err := h.postService.Create(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	var res dto.PostResponse
	if err := utils.Copy(&res, data); err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusCreated, "Success", res)
}

// Update
// @Summary		Update
// @Description	Update
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string								true	"authorization token"
// @Success		200				{object}	dto.SimpleResponse	"success"
// @Router		/v1/api/post [PUT]
func (h *PostController) Update(c echo.Context) error {
	var req dto.UpdatePostRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	userId := utils.GetUserStringIDFromContext(c)
	req.UserId = userId

	data, err := h.postService.Update(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)

	}

	var res dto.PostResponse
	if err := utils.Copy(&res, data); err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", res)
}

// GetList
// @Summary		GetList
// @Description	GetList
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string								true	"authorization token"
// @Success		200				{object}	dto.SimpleResponse	"success"
// @Router		/v1/api/post [GET]
func (h *PostController) GetList(c echo.Context) error {
	var req dto.GetListPostRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	userId := utils.GetUserStringIDFromContext(c)
	req.UserId = userId

	data, total, err := h.postService.GetList(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.ResponseList(c, "Success", total, data)
}

// Delete
// @Summary		Delete
// @Description	Delete
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string								true	"authorization token"
// @Success		200				{object}	dto.SimpleResponse	"success"
// @Router		/v1/api/post/:id [DELETE]
func (h *PostController) Delete(c echo.Context) error {
	var req dto.DeletePostRequest
	req.Id = utils.ParseStringIDFromUri(c)
	req.UserId = utils.GetUserStringIDFromContext(c)

	if err := h.postService.Delete(c.Request().Context(), req); err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", req.Id)
}

// GetOne
// @Summary		GetOne
// @Description	GetOne
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string								true	"authorization token"
// @Success		200				{object}	dto.SimpleResponse	"success"
// @Router		/v1/api/post/:id [GET]
func (h *PostController) GetOne(c echo.Context) error {
	var req dto.GetOnePostRequest
	req.Id = utils.ParseStringIDFromUri(c)

	res, err := h.postService.GetOne(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", res)
}

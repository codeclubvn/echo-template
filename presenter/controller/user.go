package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"trail_backend/pkg/utils"
	"trail_backend/presenter/request"
	"trail_backend/usecase"
)

type UserController struct {
	BaseController
	userService usecase.UserService
}

func NewUserController(userService usecase.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Update
//
//	@Summary		Update
//	@Description	Update
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string					true	"authorization token"
//	@Success		200				{object}	entity.SimpleResponse	"success"
//	@Router			/v1/api/users [PUT]
func (h *UserController) Update(c echo.Context) error {
	var req request.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	res, err := h.userService.Update(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", res)
}

// GetList
//
//	@Summary		GetList
//	@Description	GetList
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string					true	"authorization token"
//	@Success		200				{object}	entity.SimpleResponse	"success"
//	@Router			/v1/api/users [GET]
func (h *UserController) GetList(c echo.Context) error {
	var req request.GetListUserRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	data, total, err := h.userService.GetList(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.ResponseList(c, "Success", total, data)
}

// Delete
//
//	@Summary		Delete
//	@Description	Delete
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string					true	"authorization token"
//	@Success		200				{object}	entity.SimpleResponse	"success"
//	@Router			/v1/api/users/{id} [DELETE]
func (h *UserController) Delete(c echo.Context) error {
	id := utils.ParseStringIDFromUri(c)
	err := h.userService.Delete(c.Request().Context(), id)
	if err != nil {
		return h.ResponseError(c, err)
	}
	return h.Response(c, http.StatusOK, "Success", nil)
}

// GetOne
//
//	@Summary		GetOne
//	@Description	GetOne
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string					true	"authorization token"
//	@Param			id				path		string					true	"id"
//	@Success		200				{object}	entity.SimpleResponse	"success"
//	@Router			/v1/api/users/{id} [GET]
func (h *UserController) GetOne(c echo.Context) error {
	id := utils.ParseStringIDFromUri(c)
	res, err := h.userService.GetOne(c.Request().Context(), id)
	if err != nil {
		return h.ResponseError(c, err)
	}
	return h.Response(c, http.StatusOK, "Success", res)
}

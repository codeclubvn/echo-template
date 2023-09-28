package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"trial_backend/pkg/utils"
	"trial_backend/presenter/request"
	"trial_backend/usecase"
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
//	@Security		Authorization
//	@Summary		Update
//	@Description	Update
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			UpdateUserRequest	body		request.UpdateUserRequest	true	"UpdateUserRequest"
//	@Success		200					{object}	model.User					"success"
//	@Router			/v1/api/users [PUT]
func (h *UserController) Update(c echo.Context) error {
	var req request.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}
	userId, err := utils.GetUserUUIDFromContext(c)
	if err != nil {
		return h.ResponseValidationError(c, err)
	}
	req.Id = userId.String()

	res, err := h.userService.Update(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", res)
}

// Delete
//
//	@Security		Authorization
//	@Summary		Delete
//	@Description	Delete
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.SimpleResponse	"success"
//	@Router			/v1/api/users [DELETE]
func (h *UserController) Delete(c echo.Context) error {
	userId := utils.GetUserStringIDFromContext(c)
	if err := h.userService.Delete(c.Request().Context(), userId); err != nil {
		return h.ResponseError(c, err)
	}
	return h.Response(c, http.StatusOK, "Success", nil)
}

// GetOne
//
//	@Security		Authorization
//	@Summary		GetOne
//	@Description	GetOne
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.User	"success"
//	@Router			/v1/api/users [GET]
func (h *UserController) GetOne(c echo.Context) error {
	userId := utils.GetUserStringIDFromContext(c)
	res, err := h.userService.GetOne(c.Request().Context(), userId)
	if err != nil {
		return h.ResponseError(c, err)
	}
	return h.Response(c, http.StatusOK, "Success", res)
}

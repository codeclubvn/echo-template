package controller

import (
	"net/http"
	"trail_backend/dto"
	"trail_backend/utils"

	"github.com/labstack/echo/v4"
	"trail_backend/service"
)

type UserController struct {
	BaseController
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (h *UserController) Create(c echo.Context) error {
	var req dto.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	user, err := h.userService.Create(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusCreated, "Success", user.ID)
}

func (h *UserController) Update(c echo.Context) error {
	var req dto.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	userId := utils.GetUserStringIDFromContext(c)
	req.UserId = userId

	res, err := h.userService.Update(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)

	}

	return h.Response(c, http.StatusOK, "Success", res)
}

func (h *UserController) GetList(c echo.Context) error {
	var req dto.GetListUserRequest
	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	userId := utils.GetUserStringIDFromContext(c)
	req.UserId = userId

	users, err := h.userService.GetList(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", users)
}

func (h *UserController) Delete(c echo.Context) error {
	var req dto.DeleteUserRequest
	req.Id = utils.ParseStringIDFromUri(c)
	req.UserId = utils.GetUserStringIDFromContext(c)

	err := h.userService.Delete(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", nil)
}

func (h *UserController) GetOne(c echo.Context) error {
	var req dto.GetOneUserRequest
	req.Id = utils.ParseStringIDFromUri(c)

	res, err := h.userService.GetOne(c.Request().Context(), req)
	if err != nil {
		return h.ResponseError(c, err)
	}

	return h.Response(c, http.StatusOK, "Success", res)
}

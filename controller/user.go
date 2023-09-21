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

func (p *UserController) Create(c echo.Context) error {
	var req dto.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return p.ResponseValidationError(c, err)
	}

	user, err := p.userService.Create(c.Request().Context(), req)
	if err != nil {
		return p.ResponseError(c, err)
	}

	return p.Response(c, http.StatusCreated, "Success", user.ID)
}

func (p *UserController) Update(c echo.Context) error {
	var req dto.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return p.ResponseValidationError(c, err)
	}

	userID := utils.GetUserStringIDFromContext(c.Request().Context())
	_, err := p.userService.UpdateUser(c.Request().Context(), userID, req)
	if err != nil {
		return p.ResponseError(c, err)

	}

	return p.Response(c, http.StatusOK, "Success", nil)
}

func (p *UserController) List(c echo.Context) error {
	var req dto.ListUserRequest
	if err := c.Bind(&req); err != nil {
		return p.ResponseValidationError(c, err)
	}

	users, total, err := p.userService.ListUser(c.Request().Context(), req)
	if err != nil {
		return p.ResponseError(c, err)
	}

	return p.ResponseList(c, "Success", total, users)
}

func (p *UserController) Delete(c echo.Context) error {
	userID := utils.GetUserStringIDFromContext(c.Request().Context())
	err := p.userService.DeleteUser(c.Request().Context(), userID)
	if err != nil {
		return p.ResponseError(c, err)
	}

	return p.Response(c, http.StatusOK, "Success", nil)
}

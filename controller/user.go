package handler

import (
	"net/http"
	"trail_backend/utils"

	"github.com/labstack/echo/v4"
	"trail_backend/service"
)

type UserHandler struct {
	BaseController
	userService service.UserService
}

func NewUsersController(storeService service.UserService) *UserHandler {
	return &UserHandler{
		userService: storeService,
	}
}

func (p *UserHandler) CreateUser(c *echo.Context) {
	var req trail_backend.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	store, err := p.userService.CreateUserAndAssignOwner(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusCreated, "Success", store.ID)
}

func (p *UserHandler) UpdateUser(c *echo.Context) {
	var req trail_backend.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	storeID := utils.GetUserIDFromContext(c.Request.Context())
	_, err := p.userService.UpdateUser(c.Request.Context(), storeID, req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", nil)
}

func (p *UserHandler) ListUser(c *echo.Context) {
	var req trail_backend.ListUserRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	stores, total, err := p.userService.ListUser(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.ResponseList(c, "Success", total, stores)
}

func (p *UserHandler) DeleteUser(c *echo.Context) {
	storeID := utils.GetUserIDFromContext(c.Request.Context())
	err := p.userService.DeleteUser(c.Request.Context(), storeID)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Success", nil)
}

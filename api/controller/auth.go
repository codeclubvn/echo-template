package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"trail_backend/api/dto/auth"
	"trail_backend/service"

	"go.uber.org/zap"
)

type AuthController struct {
	BaseController
	authService service.AuthService
	logger      *zap.Logger
}

func NewAuthController(authService service.AuthService, logger *zap.Logger) *AuthController {
	controller := &AuthController{
		authService: authService,
		logger:      logger,
	}
	return controller
}

func (h *AuthController) Register(c echo.Context) error {
	var req dto.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	_, err := h.authService.Register(c.Request().Context(), req)

	if err != nil {
		return h.ResponseError(c, err)
	}
	return h.Response(c, http.StatusOK, "success", nil)
}

func (h *AuthController) Login(c echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	res, err := h.authService.Login(c.Request().Context(), req)

	if err != nil {
		return h.ResponseError(c, err)

	}
	return h.Response(c, http.StatusOK, "success", res)
}

package lib

import (
	"github.com/labstack/echo/v4"
	_ "trail_backend/docs"
)

type Handler struct {
	Echo *echo.Group
}

func NewServerGroupV1(instance *echo.Echo) *Handler {
	return &Handler{
		instance.Group("/v1/api"),
	}
}

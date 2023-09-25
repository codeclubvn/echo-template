package lib

import (
	"github.com/labstack/echo/v4"
	_ "trail_backend/docs"
)

type ServerGroupV1 struct {
	Echo *echo.Group
}

func NewServerGroupV1(instance *Server) *ServerGroupV1 {
	return &ServerGroupV1{
		instance.Echo.Group("/v1/api"),
	}
}

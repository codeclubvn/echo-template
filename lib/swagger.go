package lib

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "trail_backend/docs"
)

func NewSwagger(instance *echo.Echo) {
	instance.GET("/swagger/*", echoSwagger.WrapHandler)
}

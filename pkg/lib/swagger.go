package lib

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "trail_backend/presenter/docs"
)

func NewSwagger(instance *Server) {
	instance.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}

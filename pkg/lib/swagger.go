package lib

import (
	_ "echo_template/presenter/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewSwagger(instance *Server) {
	instance.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}

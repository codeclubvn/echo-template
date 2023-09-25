package route

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	"trail_backend/lib"
)

func NewSwagger(handler *lib.Handler) *lib.Handler {
	handler.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
	return handler
}

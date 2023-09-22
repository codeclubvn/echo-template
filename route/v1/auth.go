package v1

import (
	"trail_backend/api/controller"
	"trail_backend/library"
)

type AuthRoutes struct {
	handler *library.Handler
}

func NewAuthRoutes(handler *library.Handler, controller *controller.AuthController) *AuthRoutes {
	g := handler.Echo.Group("/auth")
	g.POST("/register", controller.Register)
	g.POST("/login", controller.Login)
	g.POST("/google/login", controller.GoogleLogin)
	g.POST("/google/callback", controller.GoogleCallback)
	return &AuthRoutes{
		handler: handler,
	}
}

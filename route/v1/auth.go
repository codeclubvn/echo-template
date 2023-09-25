package v1

import (
	"trail_backend/api/controller"
	"trail_backend/lib"
)

type AuthRoutes struct {
	handler *lib.Handler
}

func NewAuthRoutes(handler *lib.Handler, controller *controller.AuthController) *AuthRoutes {
	g := handler.Echo.Group("/auth")
	g.POST("/register", controller.Register)
	g.POST("/login", controller.Login)
	g.POST("/google/login", controller.GoogleLogin)
	g.POST("/google/callback", controller.GoogleCallback)
	return &AuthRoutes{
		handler: handler,
	}
}

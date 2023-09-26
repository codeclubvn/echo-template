package v1

import (
	"trail_backend/pkg/lib"
	"trail_backend/presenter/controller"
)

type AuthRoutes struct {
	handler *lib.ServerGroupV1
}

func NewAuthRoutes(handler *lib.ServerGroupV1, controller *controller.AuthController) *AuthRoutes {
	g := handler.Echo.Group("/auth")
	g.POST("/register", controller.Register)
	g.POST("/login", controller.Login)
	g.POST("/google/login", controller.GoogleLogin)
	g.POST("/google/callback", controller.GoogleCallback)
	return &AuthRoutes{
		handler: handler,
	}
}

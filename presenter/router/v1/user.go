package v1

import (
	"trial_backend/pkg/lib"
	"trial_backend/presenter/controller"
	"trial_backend/presenter/middlewares"
)

type UsersRoutes struct {
	handler *lib.ServerGroupV1
}

func NewUserRoutes(handler *lib.ServerGroupV1, c *controller.UserController, middleware *middlewares.Middleware) *UsersRoutes {
	g := handler.Echo.Group("/users")

	g.PUT("", c.Update, middleware.Auth(true))
	//g.GET("", c.GetList, middleware.Auth(true))
	g.GET("", c.GetOne, middleware.Auth(false))
	g.DELETE("", c.Delete, middleware.Auth(true))

	return &UsersRoutes{
		handler: handler,
	}
}

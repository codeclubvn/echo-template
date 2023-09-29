package v1

import (
	"echo_template/pkg/lib"
	"echo_template/presenter/controller"
	"echo_template/presenter/middlewares"
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

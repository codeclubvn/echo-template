package v1

import (
	"trail_backend/pkg/lib"
	"trail_backend/presenter/controller"
	"trail_backend/presenter/middlewares"
)

type UsersRoutes struct {
	handler *lib.ServerGroupV1
}

func NewUserRoutes(handler *lib.ServerGroupV1, c *controller.UserController, middleware *middlewares.Middleware) *UsersRoutes {
	g := handler.Echo.Group("/users")

	g.PUT("", c.Update, middleware.Auth(true))
	//g.GET("", c.GetList, middleware.Auth(true))
	g.GET("/:id", c.GetOne, middleware.Auth(false))
	g.DELETE("/:id", c.Delete, middleware.Auth(true))

	return &UsersRoutes{
		handler: handler,
	}
}

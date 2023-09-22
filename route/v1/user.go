package v1

import (
	"trail_backend/api/controller"
	"trail_backend/api/middlewares"
	"trail_backend/library"
)

type UsersRoutes struct {
	handler *library.Handler
}

func NewUserRoutes(handler *library.Handler, c *controller.UserController, middleware *middlewares.Middleware) *UsersRoutes {
	g := handler.Echo.Group("/users")

	g.POST("", c.Create, middleware.Auth(true))
	g.PUT("", c.Update, middleware.Auth(true))
	g.GET("", c.GetList, middleware.Auth(false))
	g.GET("/:id", c.GetOne, middleware.Auth(false))
	g.DELETE("/:id", c.Delete, middleware.Auth(true))

	return &UsersRoutes{
		handler: handler,
	}
}
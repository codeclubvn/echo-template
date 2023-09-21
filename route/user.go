package route

import (
	"trail_backend/controller"
	"trail_backend/library"
)

type UsersRoutes struct {
	handler *library.Handler
}

func NewUserRoutes(handler *library.Handler, c *controller.UserController) *UsersRoutes {
	g := handler.Group("/users")

	g.POST("", c.Create)
	g.PUT("", c.Update)
	g.GET("", c.List)
	g.DELETE("", c.Delete)

	return &UsersRoutes{
		handler: handler,
	}
}

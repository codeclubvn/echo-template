package v1

import (
	"trail_backend/api/controller"
	"trail_backend/library"
)

type UsersRoutes struct {
	handler *library.Handler
}

func NewUserRoutes(handler *library.Handler, c *controller.UserController) *UsersRoutes {
	g := handler.Echo.Group("/users")

	g.POST("", c.Create)
	g.PUT("", c.Update)
	g.GET("", c.GetList)
	g.GET("/:id", c.GetOne)
	g.DELETE("/:id", c.Delete)

	return &UsersRoutes{
		handler: handler,
	}
}

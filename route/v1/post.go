package v1

import (
	"trail_backend/api/controller"
	"trail_backend/api/middlewares"
	"trail_backend/library"
)

type PostRoutes struct {
	handler *library.Handler
}

func NewPostRoutes(handler *library.Handler, c *controller.PostController, middleware *middlewares.Middleware) *PostRoutes {
	g := handler.Echo.Group("/posts")

	g.POST("", c.Create, middleware.Auth(true))
	g.PUT("", c.Update, middleware.Auth(true))
	g.GET("", c.GetList, middleware.Auth(false))
	g.GET("/:id", c.GetOne, middleware.Auth(false))
	g.DELETE("/:id", c.Delete, middleware.Auth(true))

	return &PostRoutes{
		handler: handler,
	}
}

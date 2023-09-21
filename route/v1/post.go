package v1

import (
	"trail_backend/api/controller"
	"trail_backend/library"
)

type PostRoutes struct {
	handler *library.Handler
}

func NewPostRoutes(handler *library.Handler, c *controller.PostController) *PostRoutes {
	g := handler.Echo.Group("/posts")

	g.POST("", c.Create)
	g.PUT("", c.Update)
	g.GET("", c.GetList)
	g.GET("/:id", c.GetOne)
	g.DELETE("/:id", c.Delete)

	return &PostRoutes{
		handler: handler,
	}
}

package v1

import (
	"trail_backend/api/controller"
	"trail_backend/library"
)

type FileRoutes struct {
	handler *library.Handler
}

func NewFileRoutes(handler *library.Handler, c *controller.FileController) *FileRoutes {
	g := handler.Echo.Group("/files")

	g.POST("", c.Create)
	g.PUT("", c.Update)
	g.GET("", c.GetList)
	g.GET("/:id", c.GetOne)
	g.DELETE("/:id", c.Delete)

	return &FileRoutes{
		handler: handler,
	}
}

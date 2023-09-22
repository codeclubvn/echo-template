package v1

import (
	"trail_backend/api/controller"
	"trail_backend/api/middlewares"
	"trail_backend/library"
)

type FileRoutes struct {
	handler *library.Handler
}

func NewFileRoutes(handler *library.Handler, c *controller.FileController, middleware *middlewares.Middleware) *FileRoutes {
	g := handler.Echo.Group("/files")

	g.POST("", c.Create, middleware.Auth(true))
	g.PUT("", c.Update, middleware.Auth(true))
	g.GET("/:id", c.Download, middleware.Auth(false))
	g.GET("", c.GetList, middleware.Auth(false))
	g.DELETE("/:id", c.Delete, middleware.Auth(true))

	return &FileRoutes{
		handler: handler,
	}
}

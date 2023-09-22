package v1

import (
	"trail_backend/api/controller"
	"trail_backend/api/middlewares"
	"trail_backend/library"
)

type FileCloudRoutes struct {
	handler *library.Handler
}

func NewFileCloudRoutes(handler *library.Handler, c *controller.FileCloudController, middleware *middlewares.Middleware) *FileCloudRoutes {
	g := handler.Echo.Group("/image")

	g.POST("/upload", c.Upload, middleware.Auth(true))

	return &FileCloudRoutes{
		handler: handler,
	}
}

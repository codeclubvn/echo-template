package v1

import (
	"trail_backend/pkg/lib"
	"trail_backend/presenter/controller"
	"trail_backend/presenter/middlewares"
)

type FileCloudRoutes struct {
	handler *lib.ServerGroupV1
}

func NewFileCloudRoutes(handler *lib.ServerGroupV1, c *controller.FileCloudController, middleware *middlewares.Middleware) *FileCloudRoutes {
	g := handler.Echo.Group("/image")

	g.POST("/upload", c.Upload, middleware.Auth(true))

	return &FileCloudRoutes{
		handler: handler,
	}
}

package v1

import (
	"trail_backend/pkg/lib"
	"trail_backend/presenter/controller"
	"trail_backend/presenter/middlewares"
)

type FileRoutes struct {
	handler *lib.ServerGroupV1
}

func NewFileRoutes(handler *lib.ServerGroupV1, c *controller.FileController, middleware *middlewares.Middleware) *FileRoutes {
	g := handler.Echo.Group("/files")

	g.POST("", c.Upload, middleware.Auth(true))
	g.PUT("", c.Update, middleware.Auth(true))
	g.GET("/:id", c.GetOne, middleware.Auth(false))
	g.GET("/download/:id", c.Download, middleware.Auth(true))
	g.DELETE("/:id", c.Delete, middleware.Auth(true))

	return &FileRoutes{
		handler: handler,
	}
}

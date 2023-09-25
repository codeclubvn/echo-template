package v1

import (
	"trail_backend/api/controller"
	"trail_backend/api/middlewares"
	"trail_backend/lib"
)

type FileRoutes struct {
	handler *lib.Handler
}

func NewFileRoutes(handler *lib.Handler, c *controller.FileController, middleware *middlewares.Middleware) *FileRoutes {
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

package v1

import (
	"echo_template/pkg/lib"
	"echo_template/presenter/controller"
	"echo_template/presenter/middlewares"
)

type FileRoutes struct {
	handler *lib.ServerGroupV1
}

func NewFileRoutes(handler *lib.ServerGroupV1, c *controller.FileController, middleware *middlewares.Middleware) *FileRoutes {
	g := handler.Echo.Group("/files")

	g.POST("", c.SaveFile, middleware.Auth(true))
	g.PUT("", c.Update, middleware.Auth(true))
	g.GET("/:id", c.GetOne, middleware.Auth(false))
	g.GET("/download/:id", c.Download, middleware.Auth(true))
	g.DELETE("/:id", c.Delete, middleware.Auth(true))

	return &FileRoutes{
		handler: handler,
	}
}

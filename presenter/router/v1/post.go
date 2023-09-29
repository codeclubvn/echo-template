package v1

import (
	"echo_template/pkg/lib"
	"echo_template/presenter/controller"
	"echo_template/presenter/middlewares"
)

type PostRoutes struct {
	handler *lib.ServerGroupV1
}

func NewPostRoutes(handler *lib.ServerGroupV1, c *controller.PostController, middleware *middlewares.Middleware) *PostRoutes {
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

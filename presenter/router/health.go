package router

import (
	"echo_template/pkg/lib"
	"echo_template/presenter/controller"
)

type HealthRoutes struct {
	handler *lib.Server
}

func NewHealthRoutes(handler *lib.Server, c *controller.HealthController) *HealthRoutes {
	g := handler.Echo.Group("/health")
	g.GET("", c.Health)

	return &HealthRoutes{
		handler: handler,
	}
}

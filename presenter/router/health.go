package router

import (
	"trial_backend/pkg/lib"
	"trial_backend/presenter/controller"
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

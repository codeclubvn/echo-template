package route

import (
	"trail_backend/api/controller"
	"trail_backend/lib"
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

package route

import (
	"trail_backend/controller"
	"trail_backend/library"
)

type UsersRoutes struct {
	handler *library.Handler
}

func NewUserRoutes(handler *library.Handler, userController controller.UserController) *UsersRoutes {
	g := handler.Group("/users")

	g.POST("", userController.Create)
	g.PUT("", userController.Update)
	g.GET("", userController.List)
	g.DELETE("", userController.Delete)

	return &UsersRoutes{
		handler: handler,
	}
}

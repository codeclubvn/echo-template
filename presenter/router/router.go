package router

import (
	"echo_template/presenter/router/v1"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Invoke(NewHealthRoutes, v1.NewAuthRoutes, v1.NewUserRoutes, v1.NewPostRoutes, v1.NewFileRoutes, v1.NewFileCloudRoutes))

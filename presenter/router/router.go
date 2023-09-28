package router

import (
	"go.uber.org/fx"
	v12 "trial_backend/presenter/router/v1"
)

var Module = fx.Options(fx.Invoke(NewHealthRoutes, v12.NewAuthRoutes, v12.NewUserRoutes, v12.NewPostRoutes, v12.NewFileRoutes, v12.NewFileCloudRoutes))

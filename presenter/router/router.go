package router

import (
	"go.uber.org/fx"
	"trial_backend/presenter/router/v1"
)

var Module = fx.Options(fx.Invoke(NewHealthRoutes, v1.NewAuthRoutes, v1.NewUserRoutes, v1.NewPostRoutes, v1.NewFileRoutes, v1.NewFileCloudRoutes))

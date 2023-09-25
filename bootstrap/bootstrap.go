package bootstrap

import (
	"trail_backend/api/controller"
	"trail_backend/api/middlewares"
	"trail_backend/config"
	"trail_backend/infrastructure"
	"trail_backend/lib"
	"trail_backend/repository"
	"trail_backend/route"
	"trail_backend/service"
	"trail_backend/utils"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func inject() fx.Option {
	return fx.Options(
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),

		fx.Provide(
			config.NewConfig,
			utils.NewTimeoutContext,
		),
		route.Module,
		lib.Module,
		repository.Module,
		service.Module,
		controller.Module,
		middlewares.Module,
		infrastructure.Module,
	)
}

func Run() {
	fx.New(inject()).Run()
}

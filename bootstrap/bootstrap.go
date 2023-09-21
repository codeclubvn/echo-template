package bootstrap

import (
	"trail_backend/config"
	"trail_backend/controller"
	"trail_backend/infrastructure"
	"trail_backend/library"
	"trail_backend/middlewares"
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
		library.Module,
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

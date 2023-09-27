package bootstrap

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"trail_backend/config"
	"trail_backend/domain/repo"
	"trail_backend/infra"
	"trail_backend/pkg/lib"
	"trail_backend/pkg/utils"
	"trail_backend/presenter/controller"
	"trail_backend/presenter/middlewares"
	"trail_backend/presenter/router"
	"trail_backend/usecase"
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
		router.Module,
		lib.Module,
		repo.Module,
		usecase.Module,
		controller.Module,
		middlewares.Module,
		infra.Module,
	)
}

func Run() {
	fx.New(inject()).Run()
}

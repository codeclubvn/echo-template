package bootstrap

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"trial_backend/config"
	"trial_backend/domain/repo"
	"trial_backend/infra"
	"trial_backend/pkg/lib"
	"trial_backend/pkg/utils"
	"trial_backend/presenter/controller"
	"trial_backend/presenter/middlewares"
	"trial_backend/presenter/router"
	"trial_backend/usecase"
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

package bootstrap

import (
	"echo_template/config"
	"echo_template/domain/repo"
	"echo_template/infra"
	"echo_template/pkg/lib"
	"echo_template/pkg/utils"
	"echo_template/presenter/controller"
	"echo_template/presenter/middlewares"
	"echo_template/presenter/router"
	"echo_template/usecase"
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

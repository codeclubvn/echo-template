package library

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"trail_backend/api/middlewares"
	"trail_backend/config"
	"trail_backend/infrastructure"

	"go.uber.org/fx"
)

type Handler struct {
	Echo *echo.Group
}

func NewServerGroupV1(instance *echo.Echo) *Handler {
	return &Handler{
		instance.Group("/v1/api"),
	}
}

func NewServer(lifecycle fx.Lifecycle, zap *zap.Logger, config *config.Config, db *infrastructure.Database, middlewares *middlewares.Middleware) *echo.Echo {
	instance := echo.New()

	instance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	instance.Use(middleware.Logger())
	instance.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	//instance.Use(middlewares.ErrorHandler)
	// instance.Use(middlewares.JWT(config, db))

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			zap.Info("Starting HTTP server")

			go func() {
				addr := fmt.Sprint(config.Server.Host, ":", config.Server.Port)
				if err := instance.Start(addr); err != nil {
					zap.Fatal(fmt.Sprint("HTTP server failed to start %w", err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			zap.Info("Stopping HTTP server")
			return nil
		},
	})

	return instance
}

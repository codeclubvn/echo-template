package lib

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

func NewServer(lifecycle fx.Lifecycle, zap *zap.Logger, config *config.Config, db *infrastructure.Database, middlewares *middlewares.Middleware) *echo.Echo {
	instance := echo.New()

	instance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: false,
		MaxAge:           12 * 3600,
	}))
	instance.Use(middleware.Logger())
	instance.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	instance.Pre(middleware.RemoveTrailingSlash())
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

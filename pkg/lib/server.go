package lib

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
	"trail_backend/domain/repo/model"
	"trail_backend/presenter/middlewares"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"trail_backend/config"
	"trail_backend/infra"
	"trail_backend/pkg/constants"
)

type Server struct {
	Echo *echo.Echo
}

func NewServer(lifecycle fx.Lifecycle, zap *zap.Logger, config *config.Config, db *infra.Database, middlewares *middlewares.Middleware) *Server {
	instance := echo.New()
	instance.Validator = NewRequestValidator()

	instance.Use(middlewares.CORS())
	instance.Use(middleware.Logger())
	instance.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	instance.Pre(middleware.RemoveTrailingSlash())
	instance.Use(middlewares.JSONMiddleware())

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			zap.Info("Starting HTTP server")

			if err := SeedRoutes(instance, db); err != nil {
				zap.Fatal(fmt.Sprint("HTTP server failed seedRoutes %w", err))
				return nil
			}

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

	return &Server{
		instance,
	}
}

func SeedRoutes(engine *echo.Echo, db *infra.Database) error {
	// Seed permissions
	permissions := []model.Permission{}
	newPermissions := []model.Permission{}
	db.Find(&permissions)

	mapRoutes := make(map[string]model.Permission)
	for _, r := range permissions {
		mapRoutes[r.RoutePath] = r
	}

	for _, r := range engine.Routes() {
		// permission name
		// if method is GET, path is /api/v1/users, permission name is get:users
		// if method is POST, path is /api/v1/users, permission name is create:users
		// ...

		_, isPublic := constants.PublicRoutes[r.Path]
		_, isExist := mapRoutes[r.Path]

		s := strings.Split(r.Path, "/")

		if isExist {
			continue
		}

		if len(s) >= constants.NumberOfPath {
			if isPublic || s[1]+"/"+s[2] != "v1/api" {
				continue
			}
		}

		last := s[len(s)-1]
		if last == "" {
			s = s[:len(s)-1]
		}

		permissionPrefix := ""
		switch r.Method {
		case "GET":
			permissionPrefix = "get"
		case "POST":
			permissionPrefix = "create"
		case "PUT":
			permissionPrefix = "update"
		case "DELETE":
			permissionPrefix = "delete"
		}

		pn := permissionPrefix + ":" + s[len(s)-1]

		newPermissions = append(newPermissions, model.Permission{
			Method:    r.Method,
			RoutePath: r.Path,
			Name:      pn,
		})
	}

	if len(newPermissions) == 0 {
		return nil
	}

	err := db.DB.Create(&newPermissions).Error
	if err != nil {
		panic(err)
	}
	return nil
}

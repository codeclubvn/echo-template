package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (e *Middleware) CORS() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			middleware.CORSWithConfig(middleware.CORSConfig{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
				ExposeHeaders:    []string{"Content-Length", "Content-Type"},
				AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
				AllowCredentials: false,
				MaxAge:           12 * 3600,
			})
			return next(c)
		}
	}
}

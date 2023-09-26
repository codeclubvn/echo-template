package middlewares

import (
	"github.com/labstack/echo/v4"
)

func (e *Middleware) JSONMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("Content-Type", "application/json")
			return next(c)
		}
	}
}

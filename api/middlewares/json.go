package middlewares

import (
	"github.com/labstack/echo/v4"
)

func (e Middleware) JSONMiddleware(c echo.Context) {
	c.Set("Content-Type", "application/json")
}

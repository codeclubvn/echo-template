package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Api struct {
	e *echo.Echo
}

func NewApiService() *Api {
	e := echo.New()

	return &Api{
		e: e,
	}
}

func (c *Api) Run(server *http.Server) error {
	return c.e.StartServer(server)
}

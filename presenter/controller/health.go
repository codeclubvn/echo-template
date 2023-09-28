package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type HealthController struct {
	BaseController
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

// Health
//
//	@Summary		Check health server
//	@Description	Health
//	@Tags			Health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.SimpleResponse	"success"
//	@Router			/health [GET]
func (h *HealthController) Health(c echo.Context) error {
	return h.Response(c, http.StatusOK, "success", map[string]interface{}{
		"status": "UP",
		"time":   time.Now().Format("2006-01-02 15:04:05"),
		"env":    viper.GetString("server.env"),
	})
}

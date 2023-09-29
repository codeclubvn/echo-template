package middlewares

import (
	"echo_template/config"
	"echo_template/infra"
	"go.uber.org/zap"
)

type Middleware struct {
	logger *zap.Logger
	config *config.Config
	db     *infra.Database
}

func NewMiddleware(config *config.Config, db *infra.Database, logger *zap.Logger) *Middleware {
	return &Middleware{
		logger: logger,
		config: config,
		db:     db,
	}
}

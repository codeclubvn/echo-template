package middlewares

import (
	"go.uber.org/zap"
	"trial_backend/config"
	"trial_backend/infra"
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

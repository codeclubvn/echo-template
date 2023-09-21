package middlewares

import (
	"trail_backend/config"
	"trail_backend/infrastructure"

	"go.uber.org/zap"
)

type Middleware struct {
	logger *zap.Logger
	config *config.Config
	db     *infrastructure.Database
}

func NewMiddleware(config *config.Config, db *infrastructure.Database, logger *zap.Logger) *Middleware {
	return &Middleware{
		logger: logger,
		config: config,
		db:     db,
	}
}

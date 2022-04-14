package contract

import (
	"orbit_nft/db"

	"go.uber.org/zap"
)

type config struct {
	logger *zap.Logger
	sqlDB  *db.Database
}

type clientOpt func(*config)

func WithLogger(logger *zap.Logger) clientOpt {
	return func(cfg *config) {
		cfg.logger = logger
	}
}

func WithDB(sqlDB *db.Database) clientOpt {
	return func(cfg *config) {
		cfg.sqlDB = sqlDB
	}
}

func logContext(msg string) zap.Field {
	return zap.String("context", msg)
}

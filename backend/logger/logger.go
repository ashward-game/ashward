package logger

import (
	"go.uber.org/zap"
	"os"
)

var (
	defaultLogger = newWrapperLogger()
)

type wrapperLogger struct {
	base *zap.SugaredLogger
}

func newWrapperLogger() *wrapperLogger {
	var logger *zap.Logger
	switch os.Getenv("ORBIT_ENV") {
	case "production":
		logger, _ = zap.NewProduction(zap.AddCallerSkip(1))
	default:
		logger, _ = zap.NewDevelopment(zap.AddCallerSkip(1))
	}
	return &wrapperLogger{
		base: logger.Sugar(),
	}
}

// WithContext init new SugaredLogger instance with default context field(s)
func WithContext(kvs ...interface{}) *zap.SugaredLogger {
	return defaultLogger.base.With(kvs...)
}

func Shutdown() {
	_ = defaultLogger.base.Sync()
}

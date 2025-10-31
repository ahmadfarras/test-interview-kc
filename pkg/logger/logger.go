package logger

import (
	"context"

	"go.uber.org/zap"
)

// context key
type contextKey string

const loggerKey contextKey = "logger"

func WithContext(ctx context.Context, log *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, log)
}

func FromContext(ctx context.Context, defaultLogger *zap.Logger) *zap.Logger {
	if ctx == nil {
		return defaultLogger
	}
	if l, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return l
	}
	return defaultLogger
}

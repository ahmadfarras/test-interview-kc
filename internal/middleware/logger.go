package middleware

import (
	"context"
	"test-interview-kc/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func LoggerMiddleware(base *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqID := uuid.NewString()
		// create a logger with request fields
		reqLog := base.With(zap.String("request_id", reqID),
			zap.String("path", c.Path()),
			zap.String("method", c.Method()))

		// create a standard context.Context that carries the logger
		ctx := logger.WithContext(context.Background(), reqLog) // or c.UserContext() if you want to keep previous
		// attach it to fiber's user context
		c.SetUserContext(ctx)

		// log start
		reqLog.Info("Incoming request")

		// proceed to next handlers
		err := c.Next()

		// post-response log
		reqLog.Info("Request completed", zap.Int("status", c.Response().StatusCode()))

		return err
	}
}

package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap/zaptest"
)

func TestLoggerMiddleware(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"basic request"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			logger := zaptest.NewLogger(t)
			app.Use(LoggerMiddleware(logger))

			app.Get("/test", func(c *fiber.Ctx) error {
				ctx := c.UserContext()
				if ctx == nil {
					t.Error("UserContext should not be nil")
				}
				return c.SendString("ok")
			})

			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("app.Test error: %v", err)
			}
			if resp.StatusCode != fiber.StatusOK {
				t.Errorf("expected status %d, got %d", fiber.StatusOK, resp.StatusCode)
			}
		})
	}
}

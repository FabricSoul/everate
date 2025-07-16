package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// ZapRequestLogger is a middleware that logs incoming HTTP requests using a zap.Logger.
// It logs the method, URI, status, latency, remote IP, and response size.
func ZapRequestLogger(log *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			// Run the next handler in the chain
			err := next(c)
			if err != nil {
				// Let Echo's central error handler deal with it
				c.Error(err)
			}

			// After the handler has run, we can log the details
			req := c.Request()
			res := c.Response()
			stop := time.Now()

			// These are our structured logging fields
			fields := []zap.Field{
				zap.String("method", req.Method),
				zap.String("uri", req.RequestURI),
				zap.Int("status", res.Status),
				zap.String("remote_ip", c.RealIP()),
				zap.Duration("latency", stop.Sub(start)),
				zap.Int64("size", res.Size),
			}

			// Differentiate between server errors and client errors
			if res.Status >= 500 {
				log.Error("Server Error", fields...)
			} else if res.Status >= 400 {
				log.Warn("Client Error", fields...)
			} else {
				log.Info("Request Handled", fields...)
			}

			return nil
		}
	}
}

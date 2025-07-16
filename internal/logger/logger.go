// Package logger provides a configurable constructor for the application's zap logger.
// It supports distinct configurations for development and production environments.
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates a new zap.Logger tailored for either a development or production environment.
//
// In development (isDev = true), it returns a human-readable, colored, debug-level logger.
//
// In production (isDev = false), it returns a structured, JSON-formatted, info-level logger
// that is optimized for performance and machine parsing.
func New(isDev bool) (*zap.Logger, error) {
	var config zap.Config
	var err error

	if isDev {
		// Development configuration:
		// - Human-readable, console-friendly output
		// - Logs at Debug level and above
		// - Includes caller's file and line number
		// - Adds color to log levels for easy scanning
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		// Production configuration:
		// - Structured JSON output for machine parsing
		// - Logs at Info level and above
		// - Highly performant
		// - Uses ISO8601 timestamp format for consistency across services
		config = zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	// Build the logger from the selected configuration.
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

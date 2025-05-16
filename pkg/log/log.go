// Package log contains utilities for creating a new logger.
package log

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Level defines the level for logging.
type Level string

const (
	// DebugLevel is the debug level for logging.
	DebugLevel = Level("debug")

	// InfoLevel is the info level for logging.
	InfoLevel = Level("info")
)

// New returns a new logger instance.
func New(level Level) (logr.Logger, error) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var (
		zapLogLevel     zapcore.Level
		developmentMode bool
	)

	if level == DebugLevel {
		zapLogLevel = zapcore.DebugLevel
		developmentMode = true
	} else {
		zapLogLevel = zapcore.InfoLevel
	}

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapLogLevel),
		Development:      developmentMode,
		Encoding:         "console",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	zapLog, err := config.Build()
	if err != nil {
		return logr.Logger{}, fmt.Errorf("failed to build logger: %w", err)
	}

	return zapr.NewLogger(zapLog), nil
}

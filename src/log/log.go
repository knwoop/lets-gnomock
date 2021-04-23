package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates a new zap logger with the given log level, service name and environment.
func New(level zapcore.Level, serviceName, environment string) (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(level)
	config.DisableStacktrace = false
	config.InitialFields = map[string]interface{}{
		"service": serviceName,
		"env":     environment,
	}
	return config.Build()
}

// NewDiscard creates logger which output to ioutil.Discard.
// This can be used for testing.
func NewDiscard() *zap.Logger {
	return zap.NewNop()
}

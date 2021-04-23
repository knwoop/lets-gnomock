package log

import (
	"testing"

	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	l, err := New(zap.InfoLevel, "service", "config")
	if err != nil {
		t.Fatal(err)
	}
	l.Info("Info")
	l.Error("Error")
	l.Warn("Warn")
}

func TestNewDiscard(t *testing.T) {
	logger := NewDiscard()
	logger.Info("tests discard", zap.String("output", "discard"))
}

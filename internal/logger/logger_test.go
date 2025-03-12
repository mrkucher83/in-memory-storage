package logger

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func TestNewLogger(t *testing.T) {
	t.Parallel()

	logger := NewLogger()
	require.NotNil(t, logger)
	require.IsType(t, &zap.Logger{}, logger)
}

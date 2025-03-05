package compute

import (
	"github.com/mrkucher83/in-memory-storage/internal/storage"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func TestComputeExecute(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	store := storage.NewEngine(logger)
	compute := NewCompute(store, logger)

	t.Run("Valid SET command", func(t *testing.T) {
		t.Parallel()

		res, err := compute.Execute("SET key value")
		require.NoError(t, err)
		require.Equal(t, "", res)
	})

	t.Run("Valid GET command", func(t *testing.T) {
		t.Parallel()

		_, _ = compute.Execute("SET key value")
		res, err := compute.Execute("GET key")
		require.NoError(t, err)
		require.Equal(t, "value", res)
	})

	t.Run("GET non-existent key", func(t *testing.T) {
		t.Parallel()

		res, err := compute.Execute("GET missingKey")
		require.Error(t, err)
		require.Equal(t, "key not found", err.Error())
		require.Equal(t, "", res)
	})

	t.Run("Valid DEL command", func(t *testing.T) {
		t.Parallel()

		res, err := compute.Execute("DEL key")
		require.NoError(t, err)
		require.Equal(t, "", res)
	})

	t.Run("Invalid command", func(t *testing.T) {
		t.Parallel()

		res, err := compute.Execute("INVALIDCMD key value")
		require.Error(t, err)
		require.Equal(t, "invalid command syntax", err.Error())
		require.Equal(t, "", res)
	})

	t.Run("Invalid syntax", func(t *testing.T) {
		t.Parallel()

		res, err := compute.Execute("SET ключ value")
		require.Error(t, err)
		require.Equal(t, "invalid command syntax", err.Error())
		require.Equal(t, "", res)
	})
}

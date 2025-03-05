package compute

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParserParse(t *testing.T) {
	parser := &Parser{}

	t.Run("Valid SET command", func(t *testing.T) {
		t.Parallel()

		cmd, key, value := parser.Parse("SET key value")
		require.Equal(t, "SET", cmd)
		require.Equal(t, "key", key)
		require.Equal(t, "value", value)
	})

	t.Run("Valid GET command", func(t *testing.T) {
		t.Parallel()

		cmd, key, value := parser.Parse("GET key")
		require.Equal(t, "GET", cmd)
		require.Equal(t, "key", key)
		require.Empty(t, value)
	})

	t.Run("Valid DEL command", func(t *testing.T) {
		t.Parallel()

		cmd, key, value := parser.Parse("DEL key")
		require.Equal(t, "DEL", cmd)
		require.Equal(t, "key", key)
		require.Empty(t, value)
	})

	t.Run("Invalid command", func(t *testing.T) {
		t.Parallel()

		cmd, key, value := parser.Parse("INVALIDCMD key value")
		require.Equal(t, "INVALID", cmd)
		require.Empty(t, key)
		require.Empty(t, value)
	})

	t.Run("Empty input", func(t *testing.T) {
		t.Parallel()

		cmd, key, value := parser.Parse("")
		require.Equal(t, "", cmd)
		require.Empty(t, key)
		require.Empty(t, value)
	})

	t.Run("Invalid characters in input", func(t *testing.T) {
		t.Parallel()

		cmd, key, value := parser.Parse("SET ключ value")
		require.Equal(t, "INVALID", cmd)
		require.Empty(t, key)
		require.Empty(t, value)
	})
}

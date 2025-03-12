package engine

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func TestNewEngine(t *testing.T) {
	t.Parallel()

	engine := NewEngine(zap.NewNop())
	require.NotNil(t, engine)
	require.IsType(t, &Engine{}, engine)
}

func TestEngineSet(t *testing.T) {
	t.Parallel()

	engine := NewEngine(zap.NewNop())
	engine.Set("name", "Alex")
	val, ok := engine.Get("name")
	require.Equal(t, "Alex", val)
	require.True(t, ok)
}

func TestEngineGet(t *testing.T) {
	t.Parallel()

	engine := NewEngine(zap.NewNop())
	engine.Set("name", "Alex")

	t.Run("get existed value", func(t *testing.T) {
		t.Parallel()

		val, ok := engine.Get("name")
		require.Equal(t, "Alex", val)
		require.True(t, ok)
	})
	t.Run("get missing value", func(t *testing.T) {
		t.Parallel()

		val, ok := engine.Get("age")
		require.Equal(t, "", val)
		require.False(t, ok)
	})
}

func TestEngineDel(t *testing.T) {
	t.Parallel()

	engine := NewEngine(zap.NewNop())
	engine.Set("name", "Alex")
	val, ok := engine.Get("name")
	require.Equal(t, "Alex", val)
	require.True(t, ok)

	engine.Del("name")
	val, ok = engine.Get("name")
	require.False(t, ok)
	require.Empty(t, val)
}

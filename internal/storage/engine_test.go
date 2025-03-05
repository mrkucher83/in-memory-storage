package storage

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

var logger, _ = zap.NewDevelopment()

func TestNewEngine(t *testing.T) {
	t.Parallel()

	engine := NewEngine(logger)
	require.NotNil(t, engine)
	require.IsType(t, &Storage{}, engine.storage)
	require.IsType(t, &Engine{}, engine)
}

func TestEngineSetGetDel(t *testing.T) {
	t.Parallel()

	engine := NewEngine(logger)

	tests := map[string]struct {
		key  string
		val  string
		want string
	}{
		"set new key": {
			key:  "age",
			val:  "40",
			want: "40",
		},
		"set existing key": {
			key:  "name",
			val:  "John",
			want: "John",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			engine.Set(test.key, test.val)
			val, ok := engine.Get(test.key)
			require.Equal(t, test.want, val)
			require.True(t, ok)
		})
	}

	t.Run("Get missing key", func(t *testing.T) {
		val, ok := engine.Get("missingKey")
		require.False(t, ok)
		require.Empty(t, val)
	})

	t.Run("Delete key", func(t *testing.T) {
		key, value := "deleteKey", "deleteValue"
		engine.Set(key, value)
		engine.Del(key)
		_, ok := engine.Get(key)
		require.False(t, ok)
	})
}

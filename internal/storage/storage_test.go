package storage

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewStorage(t *testing.T) {
	t.Parallel()

	storage := NewStorage()
	require.NotNil(t, storage)
	require.NotNil(t, storage.data)
	require.IsType(t, &Storage{}, storage)
}

func TestStorageSet(t *testing.T) {
	t.Parallel()

	storage := &Storage{
		data: map[string]string{
			"name": "Alex",
		},
	}

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

			storage.Set(test.key, test.val)
			val, ok := storage.Get(test.key)
			require.Equal(t, test.want, val)
			require.True(t, ok)
		})
	}
}

func TestStorageGet(t *testing.T) {
	t.Parallel()

	storage := &Storage{
		data: map[string]string{
			"name": "Alex",
		},
	}

	tests := map[string]struct {
		key         string
		actualValue string
		isExist     bool
	}{
		"get missing key": {
			key:         "age",
			actualValue: "",
			isExist:     false,
		},
		"get existing key": {
			key:         "name",
			actualValue: "Alex",
			isExist:     true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			val, ok := storage.Get(test.key)
			require.Equal(t, test.actualValue, val)
			require.Equal(t, test.isExist, ok)
		})
	}
}

func TestStorageDel(t *testing.T) {
	t.Parallel()

	storage := &Storage{
		data: map[string]string{
			"name": "Alex",
		},
	}

	tests := map[string]struct {
		key string
	}{
		"delete missing key": {
			key: "age",
		},
		"delete existing key": {
			key: "name",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			storage.Del(test.key)
			val, ok := storage.Get(test.key)
			require.Equal(t, "", val)
			require.False(t, ok)
		})
	}
}

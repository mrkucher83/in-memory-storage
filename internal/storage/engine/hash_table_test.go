package engine

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewHashTable(t *testing.T) {
	t.Parallel()

	table := NewHashTable()
	require.NotNil(t, table)
	require.NotNil(t, table.data)
	require.IsType(t, &HashTable{}, table)
}

func TestHashTableSet(t *testing.T) {
	t.Parallel()

	table := &HashTable{
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

			table.Set(test.key, test.val)
			val, ok := table.Get(test.key)
			require.Equal(t, test.want, val)
			require.True(t, ok)
		})
	}
}

func TestHashTableGet(t *testing.T) {
	t.Parallel()

	table := &HashTable{
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

			val, ok := table.Get(test.key)
			require.Equal(t, test.actualValue, val)
			require.Equal(t, test.isExist, ok)
		})
	}
}

func TestHashTableDel(t *testing.T) {
	t.Parallel()

	table := &HashTable{
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

			table.Del(test.key)
			val, ok := table.Get(test.key)
			require.Equal(t, "", val)
			require.False(t, ok)
		})
	}
}

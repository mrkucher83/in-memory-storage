package compute

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestNewCompute(t *testing.T) {
	t.Parallel()

	compute := NewCompute(zap.NewNop())
	require.NotNil(t, compute)
	require.IsType(t, &Compute{}, compute)
}

func TestParse(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input       string
		expected    *Query
		expectedErr error
	}{
		"valid SET input": {
			input:       "SET name Alex",
			expected:    NewQuery("SET", []string{"name", "Alex"}),
			expectedErr: nil,
		},
		"valid GET input": {
			input:       "GET name",
			expected:    NewQuery("GET", []string{"name"}),
			expectedErr: nil,
		},
		"valid DEL input": {
			input:       "DEL name",
			expected:    NewQuery("DEL", []string{"name"}),
			expectedErr: nil,
		},
		"empty query": {
			input:       "",
			expected:    nil,
			expectedErr: ErrEmptyQuery,
		},
		"invalid SET command": {
			input:       "set name Alex",
			expected:    nil,
			expectedErr: ErrInvalidCmd,
		},
		"invalid SET number of args": {
			input:       "SET name",
			expected:    nil,
			expectedErr: ErrInvalidArgs,
		},
		"invalid GET number of args": {
			input:       "GET name Alex",
			expected:    nil,
			expectedErr: ErrInvalidArgs,
		},
		"invalid DEL number of args": {
			input:       "DEL name Alex",
			expected:    nil,
			expectedErr: ErrInvalidArgs,
		},
	}

	compute := NewCompute(zap.NewNop())
	require.NotNil(t, compute)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			query, err := compute.Parse(test.input)
			require.Equal(t, test.expectedErr, err)
			require.Equal(t, test.expected, query)
			require.True(t, reflect.DeepEqual(test.expected, query))
		})
	}
}

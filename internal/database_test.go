package database

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/mrkucher83/in-memory-storage/internal/compute"
	"github.com/mrkucher83/in-memory-storage/internal/generated/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mockStorage := mocks.NewMockstorageLayer(ctrl)
	mockCompute := mocks.NewMockcomputeLayer(ctrl)
	mockDB := NewDatabase(mockStorage, mockCompute, zap.NewNop())
	require.NotNil(t, mockDB)
	require.IsType(t, &Database{}, mockDB)
}

func TestDatabase(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mockStorage := mocks.NewMockstorageLayer(ctrl)
	mockCompute := mocks.NewMockcomputeLayer(ctrl)
	db := NewDatabase(mockStorage, mockCompute, zap.NewNop())

	t.Run("execute SET command", func(t *testing.T) {
		mockCompute.EXPECT().Parse("SET key value").Return(&compute.Query{
			Command:   "SET",
			Arguments: []string{"key", "value"},
		}, nil)
		mockStorage.EXPECT().Set("key", "value").Return()

		result, err := db.Execute("SET key value")
		require.NoError(t, err)
		require.Equal(t, "ok", result)
	})

	t.Run("execute GET command", func(t *testing.T) {
		mockCompute.EXPECT().Parse("GET key").Return(&compute.Query{Command: "GET", Arguments: []string{"key"}}, nil)
		mockStorage.EXPECT().Get("key").Return("value", nil)

		result, err := db.Execute("GET key")
		require.NoError(t, err)
		require.Equal(t, "value", result)
	})

	t.Run("execute DEL command", func(t *testing.T) {
		mockCompute.EXPECT().Parse("DEL key").Return(&compute.Query{Command: "DEL", Arguments: []string{"key"}}, nil)
		mockStorage.EXPECT().Del("key").Return()

		result, err := db.Execute("DEL key")
		require.NoError(t, err)
		require.Equal(t, "ok", result)
	})

	t.Run("execute unknown command", func(t *testing.T) {
		mockCompute.EXPECT().Parse("UNKNOWN key").Return(&compute.Query{Command: "UNKNOWN", Arguments: []string{"key"}}, nil)

		result, err := db.Execute("UNKNOWN key")
		require.Error(t, err)
		require.Equal(t, "", result)
	})

	t.Run("handle parse error", func(t *testing.T) {
		mockCompute.EXPECT().Parse("BAD QUERY").Return(nil, errors.New("parse error"))

		result, err := db.Execute("BAD QUERY")
		require.Error(t, err)
		require.Equal(t, "", result)
	})
}

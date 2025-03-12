package storage

import (
	"github.com/golang/mock/gomock"
	storageMock "github.com/mrkucher83/in-memory-storage/internal/generated/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func TestNewStorage(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mockEngine := storageMock.NewMockEngine(ctrl)
	storage := NewStorage(mockEngine, zap.NewNop())
	require.NotNil(t, storage)
	require.IsType(t, &Storage{}, storage)
}

func TestStorage_Set_Get_Del(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mockEngine := storageMock.NewMockEngine(ctrl)
	storage := NewStorage(mockEngine, zap.NewNop())

	t.Run("set and get an existing key", func(t *testing.T) {
		mockEngine.EXPECT().Set("key", "value").Return()
		mockEngine.EXPECT().Get("key").Return("value", true)

		storage.Set("key", "value")
		val, err := storage.Get("key")
		require.NoError(t, err)
		require.Equal(t, "value", val)
	})

	t.Run("get a missing key", func(t *testing.T) {
		mockEngine.EXPECT().Get("key12").Return("", false)

		val, err := storage.Get("key12")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrValNotFound)
		require.Empty(t, val)
	})

	t.Run("delete a key", func(t *testing.T) {
		mockEngine.EXPECT().Set("key5", "value5").Return()
		mockEngine.EXPECT().Get("key5").Return("value5", true)
		mockEngine.EXPECT().Del("key5").Return()
		mockEngine.EXPECT().Get("key5").Return("", false)

		storage.Set("key5", "value5")

		val, err := storage.Get("key5")
		require.NoError(t, err)
		require.Equal(t, "value5", val)

		storage.Del("key5")
		val, err = storage.Get("key5")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrValNotFound)
		require.Empty(t, val)
	})
}

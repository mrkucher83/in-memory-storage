package engine

import (
	"go.uber.org/zap"
)

type Engine struct {
	storage *HashTable
	logger  *zap.Logger
}

func NewEngine(logger *zap.Logger) *Engine {
	return &Engine{
		storage: NewHashTable(),
		logger:  logger,
	}
}

func (e *Engine) Set(key, val string) {
	e.storage.Set(key, val)
	e.logger.Debug("SET command executed", zap.String("key", key), zap.String("value", val))
}

func (e *Engine) Get(key string) (string, bool) {
	value, ok := e.storage.Get(key)
	if ok {
		e.logger.Debug("GET command executed", zap.String("key", key), zap.String("value", value))
	} else {
		e.logger.Debug("GET command - key not found", zap.String("key", key))
	}
	return value, ok
}

func (e *Engine) Del(key string) {
	e.storage.Del(key)
	e.logger.Debug("DEL command executed", zap.String("key", key))
}

package storage

import (
	"errors"
	"go.uber.org/zap"
)

var ErrValNotFound = errors.New("value not found")

type Engine interface {
	Set(string, string)
	Get(string) (string, bool)
	Del(string)
}

type Storage struct {
	engine Engine
	logger *zap.Logger
}

func NewStorage(engine Engine, logger *zap.Logger) *Storage {
	return &Storage{
		engine: engine,
		logger: logger,
	}
}

func (s *Storage) Set(key, val string) {
	s.engine.Set(key, val)
}

func (s *Storage) Get(key string) (string, error) {
	if val, ok := s.engine.Get(key); ok {
		return val, nil
	}
	return "", ErrValNotFound
}

func (s *Storage) Del(key string) {
	s.engine.Del(key)
}

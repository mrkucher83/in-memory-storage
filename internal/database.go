package database

import (
	"errors"
	"fmt"
	"github.com/mrkucher83/in-memory-storage/internal/compute"
	"go.uber.org/zap"
)

type storageLayer interface {
	Set(string, string)
	Get(string) (string, error)
	Del(string)
}

type computeLayer interface {
	Parse(string) (*compute.Query, error)
}

type Database struct {
	storageLayer storageLayer
	computeLayer computeLayer
	logger       *zap.Logger
}

func NewDatabase(storage storageLayer, compute computeLayer, logger *zap.Logger) *Database {
	return &Database{
		storageLayer: storage,
		computeLayer: compute,
		logger:       logger,
	}
}

func (d *Database) Execute(input string) (string, error) {
	d.logger.Debug("executing query", zap.String("query", input))
	query, err := d.computeLayer.Parse(input)
	if err != nil {
		return "", fmt.Errorf("failed to execute query: %w", err)
	}

	args := query.Arguments

	switch query.Command {
	case "SET":
		d.storageLayer.Set(args[0], args[1])
		return "ok", nil
	case "GET":
		val, err := d.storageLayer.Get(args[0])
		if err != nil {
			return "", fmt.Errorf("failed to get value: %w", err)
		}
		return val, nil
	case "DEL":
		d.storageLayer.Del(args[0])
		return "ok", nil
	}
	return "", errors.New("unknown server error")
}

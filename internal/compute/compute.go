package compute

import (
	"errors"
	"github.com/mrkucher83/in-memory-storage/internal/storage"
	"go.uber.org/zap"
)

type Compute struct {
	storage *storage.Engine
	parser  *Parser
	logger  *zap.Logger
}

func NewCompute(storage *storage.Engine, logger *zap.Logger) *Compute {
	return &Compute{
		storage: storage,
		parser:  &Parser{},
		logger:  logger,
	}
}

func (c *Compute) Execute(command string) (string, error) {
	cmd, key, val := c.parser.Parse(command)

	if cmd == "INVALID" {
		c.logger.Debug("Invalid command syntax", zap.String("command", command))
		return "", errors.New("invalid command syntax")
	}

	switch cmd {
	case "SET":
		c.storage.Set(key, val)
		return "", nil
	case "GET":
		if value, ok := c.storage.Get(key); ok {
			return value, nil
		}
		return "", errors.New("key not found")
	case "DEL":
		c.storage.Del(key)
		return "", nil
	default:
		return "", errors.New("unknown command")
	}
}

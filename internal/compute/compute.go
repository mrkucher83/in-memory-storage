package compute

import (
	"errors"
	"go.uber.org/zap"
	"strings"
)

var (
	ErrEmptyQuery  = errors.New("empty query")
	ErrInvalidCmd  = errors.New("invalid command")
	ErrInvalidArgs = errors.New("invalid arguments")
)

type Query struct {
	Command   string
	Arguments []string
}

func NewQuery(cmd string, args []string) *Query {
	return &Query{
		Command:   cmd,
		Arguments: args,
	}
}

type Compute struct {
	logger *zap.Logger
}

func NewCompute(logger *zap.Logger) *Compute {
	return &Compute{logger: logger}
}

func (c *Compute) Parse(input string) (*Query, error) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil, ErrEmptyQuery
	}

	cmd := parts[0]
	if !isValidCommand(cmd) {
		c.logger.Debug("invalid command", zap.String("query", cmd))
		return nil, ErrInvalidCmd
	}

	if !isValidArgsCount(cmd, parts[1:]) {
		c.logger.Debug("invalid arguments", zap.String("arguments", input))
		return nil, ErrInvalidArgs
	}

	query := NewQuery(cmd, parts[1:])
	return query, nil
}

func isValidCommand(cmd string) bool {
	switch cmd {
	case "SET", "GET", "DEL":
		return true
	default:
		return false
	}
}

func isValidArgsCount(cmd string, args []string) bool {
	switch cmd {
	case "SET":
		if len(args) == 2 {
			return true
		}
	case "GET", "DEL":
		if len(args) == 1 {
			return true
		}
	}

	return false
}

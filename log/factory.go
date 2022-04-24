package log

import (
	"sync"

	"github.com/jacksonCLyu/ridi-faces/pkg/logger"
)

var logMemCache sync.Map

// GetLogger returns a logger with the given name.
func GetLogger(name string) logger.Logger {
	l, _ := logMemCache.LoadOrStore(name, NewLogger(WithName(name)))
	return l.(logger.Logger)
}

// GetLoggerWithOptions returns a logger with the given name and options.
func GetLoggerWithOptions(name string, opts ...Option) logger.Logger {
	l, _ := logMemCache.LoadOrStore(name, NewLogger(opts...))
	return l.(logger.Logger)
}
